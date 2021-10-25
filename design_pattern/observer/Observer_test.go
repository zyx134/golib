package observer

import "testing"

func TestObserver(t *testing.T) {
	sub := &Subject{}
	a := &AObserver{Id: "A"}
	b := &AObserver{Id: "b"}
	sub.Attach(a, b)
	sub.SetState("hello world")
	sub.SetState("i know")
}
