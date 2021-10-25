package _filepath

import (
	"fmt"
	"os"
	"testing"
)

func TestSplit(t *testing.T) {
	fmt.Println(Split("/home/polaris/studygolang"))
}

func TestJoin(t *testing.T) {
	fmt.Println(Join("a", "b", "c", "dd"))
}
func TestSplitList(t *testing.T) {
	fmt.Println(SplitList("/home/polaris/studygolang"))
}
func TestClean(t *testing.T) {
	fmt.Println(Clean("home1/pola2ris/studygol3ang"))
}

func TestWalk(t *testing.T) {
	fp := "/Users/zyx/go/lib"
	Walk(fp, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
