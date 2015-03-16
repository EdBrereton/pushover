package pushover

// AssembleError represents any errors occurring within the
// pushover message generation
type AssembleError int

const (
	ErrNoError AssembleError = iota
	ErrNoToken
	ErrNoUser
	ErrNoMsg
	ErrNoRetry
	ErrNoExpire
	ErrSendFail
	ErrJsonFail
	ErrMsgTooLong
	ErrDeviceTooLong
	ErrTitleTooLong
	ErrUrlTitleTooLong
	ErrUrlTooLong
	ErrCallbackTooLong
	ErrInvalidToken
	ErrInvalidUser
	ErrInvalidRetry
	ErrInvalidExpire
)

// PushoverPriority represents the priority applied to the message
type PushoverPriority int

const (
	PpLowest PushoverPriority = iota - 2
	PpLow
	PpNormal
	PpHigh
	PpEmergency
)

// PushoverSound represents the audio alert to be used for the alert
type PushoverSound int

const (
	PsDefault PushoverSound = iota
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

// String is used to convert PushoverSound to an apropriate value
// for pushover.net
func (p PushoverSound) String() string {
	sounds := [...]string{"", "pushover", "bike", "bugle", "cashregister",
		"classical", "cosmic", "falling", "gamelan", "incoming", "intermission",
		"magic", "mechanical", "pianobar", "siren", "spacealarm", "tugboat",
		"alien", "climb", "persistent", "echo", "updown", "none"}

	return sounds[p]
}
