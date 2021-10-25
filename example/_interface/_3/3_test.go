package _3

import "testing"

func Test(t *testing.T) {
	var w WashingMachine
	h := haier{dryer{}}
	w = h
	w.dry()
	w.wash()
}
