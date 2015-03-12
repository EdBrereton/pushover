package pushover

import "testing"
import "net/url"

func (m *PushOverMessage) TestAssemble() (msg url.Values, err AssembleError) {
	msg, err = m.assemble()
	return
}

func TestMessageAssemble_NoToken(t *testing.T) {

	message := PushOverMessage{}
	message.User = "PUSHOVER_USER"
	message.Message = "Message"

	_, err := message.TestAssemble()
	if err != ErrNoToken {
		t.Error("Token not checked correctly")
	}
}

func TestMessageAssemble_NoUser(t *testing.T) {

	message := PushOverMessage{}
	message.Token = "PUSHOVER_TOKEN"
	message.Message = "Message"

	_, err := message.TestAssemble()
	if err != ErrNoUser {
		t.Error("User not checked correctly")
	}
}

func TestMessageAssemble_NoMessage(t *testing.T) {

	message := PushOverMessage{}
	message.Token = "PUSHOVER_TOKEN"
	message.User = "PUSHOVER_USER"

	_, err := message.TestAssemble()
	if err != ErrNoMsg {
		t.Error("Message not checked correctly")
	}
}

func TestMessageAssemble_MinimalMessage(t *testing.T) {

	message := PushOverMessage{}
	message.Token = "PUSHOVER_TOKEN"
	message.User = "PUSHOVER_USER"
	message.Message = "Message"

	vals, err := message.TestAssemble()
	if err != ErrNoError {
		t.Error("Minimal Message not assembled correctly")
	}

	if vals.Get("token") != "PUSHOVER_TOKEN" {
		t.Error("Token not set correctly")
	}

	if vals.Get("user") != "PUSHOVER_USER" {
		t.Error("User not set correctly")
	}

	if vals.Get("message") != "Message" {
		t.Error("Message not set correctly")
	}
}

func TestMessageAssemble_FullMessage(t *testing.T) {

	message := PushOverMessage{}
	message.Token = "PUSHOVER_TOKEN"
	message.User = "PUSHOVER_USER"
	message.Message = "Message"
	message.Device = "Device"
	message.Priority = PpHigh
	message.Title = "Title"
	message.Sound = PsCosmic
	message.Url, _ = url.Parse("http://www.google.com")
	message.Url_title = "Url_Title"
	message.Timestamp = "Timestamp"

	vals, err := message.TestAssemble()
	if err != ErrNoError {
		t.Error("Minimal Message not assembled correctly")
	}

	if vals.Get("token") != "PUSHOVER_TOKEN" {
		t.Error("Token not set correctly")
	}

	if vals.Get("user") != "PUSHOVER_USER" {
		t.Error("User not set correctly")
	}

	if vals.Get("message") != "Message" {
		t.Error("Message not set correctly")
	}

	if vals.Get("device") != "Device" {
		t.Error("Device not set correctly")
	}

	if vals.Get("title") != "Title" {
		t.Error("Title not set correctly")
	}

	if vals.Get("url") != "http://www.google.com" {
		t.Error("Url not set correctly")
	}

	if vals.Get("url_title") != "Url_Title" {
		t.Error("Url_Title not set correctly")
	}

	if vals.Get("priority") != "1" {
		t.Errorf("Priority not set correctly")
	}

	if vals.Get("timestamp") != "Timestamp" {
		t.Error("Timestamp not set correctly")
	}

	if vals.Get("sound") != "cosmic" {
		t.Error("Sound not set correctly")
	}
}
