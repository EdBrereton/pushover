package pushover

import "testing"

func TestPushoverSound(t *testing.T) {
	cases := []struct {
		in   PushoverSound
		want string
	}{
		{PsDefault, ""},
		{PsPushover, "pushover"},
		{PsBike, "bike"},
		{PsBugle, "bugle"},
		{PsCashregister, "cashregister"},
		{PsClassical, "classical"},
		{PsCosmic, "cosmic"},
		{PsFalling, "falling"},
		{PsGamelan, "gamelan"},
		{PsIncoming, "incoming"},
		{PsIntermission, "intermission"},
		{PsMagic, "magic"},
		{PsMechanical, "mechanical"},
		{PsPianobar, "pianobar"},
		{PsSiren, "siren"},
		{PsSpacealarm, "spacealarm"},
		{PsTugboat, "tugboat"},
		{PsAlien, "alien"},
		{PsClimb, "climb"},
		{PsPersistent, "persistent"},
		{PsEcho, "echo"},
		{PsUpdown, "updown"},
		{PsNone, "none"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
