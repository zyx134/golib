package _type

import (
	"testing"
)

func TestJustifyType(t *testing.T) {
	var x interface{}
	x = "abc"
	JustifyType(x)
}
func TestJustifyType2(t *testing.T) {
	var x interface{} = 2
	JustifyType(x)
}
func TestJustifyType3(t *testing.T) {
	var x interface{}
	a := true
	x = a
	JustifyType(x)
}
func TestGetJi(t *testing.T) {
	GetJi()
}
