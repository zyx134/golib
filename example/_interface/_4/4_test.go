package _4

import "testing"

func Test(t *testing.T) {
	var a animal
	c := cat{
		name: "猪猪莫",
	}
	a = c
	a.move()
	a.say()
}
