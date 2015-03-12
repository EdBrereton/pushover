package pushover

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
