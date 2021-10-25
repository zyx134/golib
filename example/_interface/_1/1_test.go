package _1

import "testing"

func Test(t *testing.T) {
	var x Sayer
	var y Mover

	var a = dog{
		name: "ddd",
	}
	x = a
	y = a
	x.say()
	y.move()
}
