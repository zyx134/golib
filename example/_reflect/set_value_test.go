package _reflect

import "testing"

func TestSetValue(t *testing.T) {
	SetValue(1)
	a := 2.2
	SetValue(&a)
}
