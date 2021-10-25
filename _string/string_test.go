package _string

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	count := Count("asdasdasdqwe1dfgdfg", "a")
	fmt.Printf("出现的次数%v\n", count)

}
func TestChineseToPinyin(t *testing.T) {
	s, err := ChineseToPinyin("阿叔更好的", "")
	fmt.Printf("%v,%v\n", s, err)

}
