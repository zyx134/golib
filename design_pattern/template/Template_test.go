package template

import "testing"

func TestTemplate1(t *testing.T) {
	RunGame(&FootBall{})
	RunGame(&Cricket{})
}
