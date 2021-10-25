package _string

import (
	"fmt"
	"testing"
)

func TestParseInt(t *testing.T) {
	n, err := ParseInt("128", 10, 8)
	fmt.Printf("%v,  %v \n", n, err)
}
