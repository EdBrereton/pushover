// A package to interface with the Pushover.net service
package pushover

import "net/url"
import "net/http"
import "fmt"
import "encoding/json"
import "io/ioutil"

// PushOverMessage is the basic message type used to construct messages
// to send via the pushover service
type PushOverMessage struct {
	Token     string
	User      string
	Message   string
	Device    string
	Title     string
	Url       *url.URL
	Url_title string
	Priority  PushoverPriority
	Timestamp string
	Sound     PushoverSound
}

// Send is the method used to send the generated message.
// The response from the service is returned, along with an error code
func (m *PushOverMessage) Send() (reply Response, err AssembleError) {
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
		err = ErrJsonFail
		return
	}

	return
}

// Assemble is used to generate the URL values from the populated pushovermessage
func (m *PushOverMessage) assemble() (msg url.Values, err AssembleError) {
	if m.Token == "" {
		err = ErrNoToken
	}

	if m.User == "" {
		err = ErrNoUser
	}

	if m.Message == "" {
		err = ErrNoMsg
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

	if m.Url.String() != "" {
		msg.Add("url", m.Url.String())
	}

	if m.Url_title != "" {
		msg.Add("url_title", m.Url_title)
	}

	if m.Priority != PpNormal {
		msg.Add("priority", fmt.Sprintf("%d", m.Priority))
	}

	if m.Timestamp != "" {
		msg.Add("timestamp", m.Timestamp)
	}

	if m.Sound != PsDefault {
		msg.Add("sound", m.Sound.String())
	}

	return
}
