package state

import "testing"

func TestState(t *testing.T) {
	l := Light{State: &OnLightState{}}
	l.PressSwitch()
	l.PressSwitch()
	l.PressSwitch()
	l.PressSwitch()
}
