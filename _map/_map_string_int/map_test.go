package _map_string_int

import (
	"fmt"
	"testing"
)

func TestDelValue(t *testing.T) {
	fmt.Println(DelValue(map[string]int{
		"a": 1, "b": 2,
	}, "a"))
}

func TestClear(t *testing.T) {
	fmt.Println(Clear(map[string]int{"a": 1}))
}

func TestKeyExist(t *testing.T) {
	fmt.Println(KeyExist(map[string]int{"a": 1}, "a"))
}
