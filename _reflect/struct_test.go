package _reflect

import (
	"fmt"
	"testing"
)

//结构体字段必须是大写
type test struct {
	A string
	B string
	c int
}

func TestStructToMapDemo(t *testing.T) {
	student := test{"aa", "bb", 18}
	data := Struct2Map(student)
	fmt.Println(data)
}
