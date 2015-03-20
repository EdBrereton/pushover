// Package pushover is designed to interface with the Pushover.net service
// by superblock
package pushover

import "net/url"
import "net/http"
import "fmt"
import "encoding/json"
import "io/ioutil"
import "time"

// Message is the basic message type used to construct messages
// to send via the pushover service. The token, user and message are always
// mandatory, while retry and expire are mandatory is the message priority is
// PpEmergency. Missing a mandatory value will result in an error when trying
// to send
type Message struct {
	Token     string
	User      string
	Message   string
	Device    string
	Title     string
	URL       *url.URL
	URLTitle  string
	Priority  Priority
	Timestamp time.Time
	Sound     Sound
	Retry     int
	Expire    int
	HTML      bool
	Callback  *url.URL
}

// Send is the method used to send the generated message.
// The response from the service is returned, along with an error code to
// indicate any issues encountered
func (m *Message) Send() (reply Response, err AssembleError) {
	val, err := m.assemble()
	if err != ErrNoError {
		return
	}

	response, postErr := http.PostForm("https://api.pushover.net/1/messages.json", val)
	if postErr != nil {
		err = ErrSendFail
		return
	}

	defer response.Body.Close()

	data, postErr := ioutil.ReadAll(response.Body)
	if postErr != nil {
		err = ErrSendFail
		return
	}

	postErr = json.Unmarshal(data, &reply)
	if postErr != nil {
		err = ErrJSONFail
		return
	}

	return
}

// CheckLengths checks the lengths of any fields in the message which have
// a maximum length
func (m *Message) checkLengths() (err AssembleError) {
	if len(m.Message) > 1024 {
		err = ErrMsgTooLong
		return
	}

	if len(m.Device) > 25 {
		err = ErrDeviceTooLong
		return
	}

	if len(m.Title) > 250 {
		err = ErrTitleTooLong
		return
	}

	if len(m.URLTitle) > 100 {
		err = ErrURLTitleTooLong
		return
	}

	if m.URL != nil {
		if len(m.URL.String()) > 512 {
			err = ErrURLTooLong
			return
		}
	}

	if m.Callback != nil {
		if len(m.Callback.String()) > 512 {
			err = ErrCallbackTooLong
			return
		}
	}

	return
}

// CheckValid performs some basic checks on values to ensure that basic
// errors like missing characters in the token are caught before sending
func (m *Message) checkValid() (err AssembleError) {
	if len(m.Token) != 30 {
		err = ErrInvalidToken
		return
	}

	if len(m.User) != 30 {
		err = ErrInvalidUser
		return
	}

	if m.Retry != 0 && m.Retry < 30 {
		err = ErrInvalidRetry
		return
	}

	if m.Expire > 86400 {
		err = ErrInvalidExpire
		return
	}

	return
}

// CheckMandatory ensures that all fields that are required are present
func (m *Message) checkMandatory() (err AssembleError) {
	if m.Token == "" {
		err = ErrNoToken
		return
	}

	if m.User == "" {
		err = ErrNoUser
		return
	}

	if m.Message == "" {
		err = ErrNoMsg
		return
	}

	if m.Priority == PpEmergency {
		if m.Retry == 0 {
			err = ErrNoRetry
			return
		}

		if m.Expire == 0 {
			err = ErrNoExpire
			return
		}

	}

	return
}

// Validate is used to validate that required fields are filled, and all fields are
// within acceptable ranges
func (m *Message) validate() (err AssembleError) {

	err = m.checkMandatory()
	if err != ErrNoError {
		return
	}

	err = m.checkLengths()
	if err != ErrNoError {
		return
	}

	err = m.checkValid()
	if err != ErrNoError {
		return
	}

	return

}

// Assemble is used to generate the URL values from the populated Message
func (m *Message) assemble() (msg url.Values, err AssembleError) {

	err = m.validate()
	if err != ErrNoError {
		return
	}

	msg = url.Values{}
	msg.Set("token", m.Token)
	msg.Add("user", m.User)
	msg.Add("message", m.Message)

	if m.Device != "" {
		msg.Add("device", m.Device)
	}

	if m.Title != "" {
		msg.Add("title", m.Title)
	}

	if m.URL != nil {
		if m.URL.String() != "" {
			msg.Add("url", m.URL.String())
		}
	}

	if m.Callback != nil {
		if m.Callback.String() != "" {
			msg.Add("callback", m.Callback.String())
		}
	}

	if m.URLTitle != "" {
		msg.Add("url_title", m.URLTitle)
	}

	if m.Priority != PpNormal {
		msg.Add("priority", fmt.Sprintf("%d", m.Priority))
	}

	if !m.Timestamp.IsZero() {
		msg.Add("timestamp", fmt.Sprintf("%d", m.Timestamp.Unix()))
	}

	if m.Sound != PsDefault {
		msg.Add("sound", m.Sound.String())
	}

	return
}
