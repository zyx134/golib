package _reflect

import "testing"

func TestGetValue(t *testing.T) {
	GetValue(123)
	GetValue("bbb")
	GetValue(true)
	GetValue(123.222)
	GetValue([]string{"a", "b", "c"})
	GetValue(map[int]int{1: 1})
	GetValue(func(a int) {})
}
