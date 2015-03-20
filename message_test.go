package pushover

import "testing"
import "net/url"
import "time"

const (
	TestUser  = "987654321098765432109876543210"
	TestToken = "012345678901234567890123456789"
)

func (m *Message) TestAssemble() (msg url.Values, err AssembleError) {
	msg, err = m.assemble()
	return
}

func (m *Message) TestValidate() (err AssembleError) {
	err = m.validate()
	return
}

func (m *Message) TestCheckMandatory() (err AssembleError) {
	err = m.checkMandatory()
	return
}

func (m *Message) TestCheckLengths() (err AssembleError) {
	err = m.checkLengths()
	return
}

func (m *Message) TestCheckValid() (err AssembleError) {
	err = m.checkValid()
	return
}

func TestMessageCheckLengths_Message(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = "TestUser"

	for i := 0; i < 1025; i++ {
		message.Message = message.Message + "X"
	}

	err := message.TestCheckLengths()
	if err != ErrMsgTooLong {
		t.Error("Message length not checked correctly")
	}
}

func TestMessageCheckLengths_Device(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"

	for i := 0; i < 31; i++ {
		message.Device = message.Device + "X"
	}

	err := message.TestCheckLengths()
	if err != ErrDeviceTooLong {
		t.Error("Device length not checked correctly")
	}
}

func TestMessageCheckLengths_Title(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"

	for i := 0; i < 251; i++ {
		message.Title = message.Title + "X"
	}

	err := message.TestCheckLengths()
	if err != ErrTitleTooLong {
		t.Error("Title length not checked correctly")
	}
}

func TestMessageCheckLengths_UrlTitle(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"

	for i := 0; i < 101; i++ {
		message.URLTitle = message.URLTitle + "X"
	}

	err := message.TestCheckLengths()
	if err != ErrURLTitleTooLong {
		t.Error("URL Title length not checked correctly")
	}
}

func TestMessageCheckLengths_Url(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"

	var address string
	address = "http://www."
	for i := 0; i < 500; i++ {
		address = address + "X"
	}
	address = address + ".com"

	message.URL, _ = url.Parse(address)

	err := message.TestCheckLengths()
	if err != ErrURLTooLong {
		t.Error("URL length not checked correctly")
	}
}

func TestMessageCheckLengths_Callback(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"

	var address string
	address = "http://www."
	for i := 0; i < 500; i++ {
		address = address + "X"
	}
	address = address + ".com"

	message.Callback, _ = url.Parse(address)

	err := message.TestCheckLengths()
	if err != ErrCallbackTooLong {
		t.Error("Callback length not checked correctly")
	}
}

func TestMessageCheckValid_Token(t *testing.T) {

	message := Message{}
	message.Token = "0123456789012345678901234567890"
	message.User = TestUser
	message.Message = "Message"

	err := message.TestCheckValid()
	if err != ErrInvalidToken {
		t.Error("Token length not checked correctly")
	}
}

func TestMessageCheckValid_User(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = "0123456789012345678901234567890"
	message.Message = "Message"

	err := message.TestCheckValid()
	if err != ErrInvalidUser {
		t.Error("User length not checked correctly")
	}
}

func TestMessageCheckValid_Retry(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"
	message.Retry = 15

	err := message.TestCheckValid()
	if err != ErrInvalidRetry {
		t.Error("Retry value not checked correctly")
	}
}

func TestMessageCheckValid_Expire(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"
	message.Expire = 86401

	err := message.TestCheckValid()
	if err != ErrInvalidExpire {
		t.Error("Expiry not checked correctly")
	}
}

func TestMessageCheckMandatory_Token(t *testing.T) {

	message := Message{}
	message.User = TestUser
	message.Message = "Message"

	err := message.TestCheckMandatory()
	if err != ErrNoToken {
		t.Error("Token not checked correctly")
	}
}

func TestMessageCheckMandatory_User(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.Message = "Message"

	err := message.TestCheckMandatory()
	if err != ErrNoUser {
		t.Error("User not checked correctly")
	}
}

func TestMessageCheckMandatory_Message(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser

	err := message.TestCheckMandatory()
	if err != ErrNoMsg {
		t.Error("Message not checked correctly")
	}
}

func TestMessageCheckMandatory_Retry(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"
	message.Priority = PpEmergency
	message.Expire = 300

	err := message.TestCheckMandatory()
	if err != ErrNoRetry {
		t.Error("Retry not checked correctly")
	}
}

func TestMessageCheckMandatory_Expire(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"
	message.Priority = PpEmergency
	message.Retry = 35

	err := message.TestCheckMandatory()
	if err != ErrNoExpire {
		t.Error("Expiry not checked correctly")
	}
}

func TestMessageAssemble_MandatoryError(t *testing.T) {

	message := Message{}
	message.User = TestUser
	message.Message = "Message"

	_, err := message.TestAssemble()
	if err == ErrNoError {
		t.Error("Minimal Message not checked on assembly correctly (checkMandatory)")
	}
}

func TestMessageAssemble_LengthError(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"

	var address string
	address = "http://www."
	for i := 0; i < 500; i++ {
		address = address + "X"
	}
	address = address + ".com"

	message.URL, _ = url.Parse(address)

	_, err := message.TestAssemble()
	if err == ErrNoError {
		t.Error("Minimal Message not checked on assembly correctly (checkLengths)")
	}
}

func TestMessageAssemble_VadliityError(t *testing.T) {

	message := Message{}
	message.Token = "0123456789012345678901234567890"
	message.User = TestUser
	message.Message = "Message"

	_, err := message.TestAssemble()
	if err == ErrNoError {
		t.Error("Minimal Message not checked on assembly correctly (checkValid)")
	}
}

func TestMessageAssemble_MinimalMessage(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"

	vals, err := message.TestAssemble()
	if err != ErrNoError {
		t.Error("Minimal Message not assembled correctly")
	}

	if vals.Get("token") != TestToken {
		t.Error("Token not set correctly")
	}

	if vals.Get("user") != TestUser {
		t.Error("User not set correctly")
	}

	if vals.Get("message") != "Message" {
		t.Error("Message not set correctly")
	}
}

func TestMessageAssemble_FullMessage(t *testing.T) {

	message := Message{}
	message.Token = TestToken
	message.User = TestUser
	message.Message = "Message"
	message.Device = "Device"
	message.Priority = PpHigh
	message.Title = "Title"
	message.Sound = PsCosmic
	message.Callback, _ = url.Parse("http://www.callback.com")
	message.URL, _ = url.Parse("http://www.google.com")
	message.URLTitle = "Url_Title"
	message.Timestamp = time.Unix(60, 0)

	vals, err := message.TestAssemble()
	if err != ErrNoError {
		t.Error("Full Message not assembled correctly")
	}

	if vals.Get("token") != TestToken {
		t.Error("Token not set correctly")
	}

	if vals.Get("user") != TestUser {
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

	if vals.Get("callback") != "http://www.callback.com" {
		t.Error("Callback not set correctly")
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

	if vals.Get("timestamp") != "60" {
		t.Error("Timestamp not set correctly")
	}

	if vals.Get("sound") != "cosmic" {
		t.Error("Sound not set correctly")
	}
}
