package _reflect

import (
	"fmt"
	"testing"
)

func TestGetType(t *testing.T) {
	GetType(222)
	GetType(222.1)
	GetType(true)
	GetType([]int{1, 2, 3})
	GetType(map[int]int{1: 1})
	GetType(func(a int) { fmt.Println(a) })
}
