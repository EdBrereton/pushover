package pushover

import "net/url"
import "net/http"
import "fmt"

type PushOverMessage struct {
	Token     string
	User      string
	Message   string
	Device    string
	Title     string
	Url       url.URL
	Url_title string
	Priority  int
	Timestamp string
	Sound     PushoverSound
}

func (m *PushOverMessage) Send() (err AssembleError) {
	val, err := m.assemble()
	if err != errNoError {
		return
	}

	http.PostForm("https://api.pushover.net/1/messages.json", val)

	return
}

func (m *PushOverMessage) assemble() (msg url.Values, err AssembleError) {
	if m.Token == "" {
		err = errNoToken
	}

	if m.User == "" {
		err = errNoUser
	}

	if m.Message == "" {
		err = errNoMsg
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

	if m.Priority != 0 {
		msg.Add("priority", fmt.Sprintf("%d", m.Priority))
	}

	if m.Timestamp != "" {
		msg.Add("timestamp", m.Timestamp)
	}

	if m.Sound != PsPushover {
		msg.Add("sound", m.Sound.String())
	}

	return
}
