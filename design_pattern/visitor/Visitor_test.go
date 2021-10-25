package visitor

import "testing"

func TestVisitor(t *testing.T) {
	d := ABData{A: 111, B: 222}
	add := &AddVisitor{}
	sub := &SubVisitor{}
	d.Accept(add)
	d.Accept(sub)
}
