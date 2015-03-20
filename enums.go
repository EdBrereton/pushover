package pushover

// AssembleError represents any errors occurring within the
// pushover message generation
type AssembleError int

// These are the valid values of AssembleError
const (
	ErrNoError AssembleError = iota
	ErrNoToken
	ErrNoUser
	ErrNoMsg
	ErrNoRetry
	ErrNoExpire
	ErrSendFail
	ErrJSONFail
	ErrMsgTooLong
	ErrDeviceTooLong
	ErrTitleTooLong
	ErrURLTitleTooLong
	ErrURLTooLong
	ErrCallbackTooLong
	ErrInvalidToken
	ErrInvalidUser
	ErrInvalidRetry
	ErrInvalidExpire
)

// Priority represents the priority applied to the message
type Priority int

// These are the valid values of Priority
const (
	PpLowest Priority = iota - 2
	PpLow
	PpNormal
	PpHigh
	PpEmergency
)

// Sound represents the audio alert to be used for the alert
type Sound int

//These are the valid values of Sound
const (
	PsDefault Sound = iota
	PsPushover
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

// String is used to convert Sound to an apropriate value
// for pushover.net
func (p Sound) String() string {
	sounds := [...]string{"", "pushover", "bike", "bugle", "cashregister",
		"classical", "cosmic", "falling", "gamelan", "incoming", "intermission",
		"magic", "mechanical", "pianobar", "siren", "spacealarm", "tugboat",
		"alien", "climb", "persistent", "echo", "updown", "none"}

	return sounds[p]
}
