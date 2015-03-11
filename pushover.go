package pushover

import "net/url"
import "net/http"
import "fmt"

type AssembleError int

const (
	errNoError AssembleError = 0 << iota
	errNoToken
	errNoUser
	errNoMsg
)

type PushoverSound int

const (
	PsPushover PushoverSound = 0 << iota
	PsBike
	PsBugle
	PsCashregister
	PsClassical
	PsCosmic
	PsFalling
	PsGamelan
	PsIncoming
	PsIntermission
	PsMagic
	PsMechanical
	PsPianobar
	PsSiren
	PsSpacealarm
	PsTugboat
	PsAlien
	PsClimb
	PsPersistent
	PsEcho
	PsUpdown
	PsNone
)

func (p PushoverSound) String() string {
	sounds := [...]string{"pushover", "bike", "bugle", "cashregister",
		"classical", "cosmic", "falling", "gamelan", "incoming", "intermission",
		"magic", "mechanical", "pianobar", "siren", "spacealarm", "tugboat",
		"alien", "climb", "persistent", "echo", "updown", "none"}

	return sounds[p]
}

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
