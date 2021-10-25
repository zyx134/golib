package _2

import "testing"

func Test(t *testing.T) {
	var x Mover
	a := dog{name: "dd"}
	b := car{name: "宝马"}
	x = a
	x.move()
	x = b
	x.move()
}
