package _sort

import (
	"fmt"
	"testing"
)

func TestFileSort(t *testing.T) {
	list := []string{
		"1.txt",
		"100.abc",
		"101.txy",
		"2.txt",
		"333.txy",
	}
	fmt.Println(FileSort(list))
}
