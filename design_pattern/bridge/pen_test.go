package bridge

import "testing"

func TestPen(t *testing.T) {
	var c Color
	var r = new(Red)
	c = r
	NewBrush("big", c).Draw()
	var g = new(Green)
	c = g
	NewBrush("big", c).Draw()
	var y = new(Yellow)
	c = y
	NewBrush("small", c).Draw()
}
