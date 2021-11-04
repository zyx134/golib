package copy

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	a := []string{"a"}
	var b []string
	err := Copy(a, &b)
	fmt.Println(err)
	a = append(a, "b")
	fmt.Println("a", a)
	fmt.Println("b", b)
}
func TestCopy1(t *testing.T) {
	a := make(map[string]string)
	b := make(map[string]string)
	a["a"] = "a"
	err := Copy(a, &b)
	a["b"] = "b"
	fmt.Println(err)
	fmt.Println("a", a)
	fmt.Println("b", b)
}
