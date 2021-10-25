package _reflect

import (
	"fmt"
	"reflect"
)

func SetValueSlice(a []int) {
	v := reflect.ValueOf(a)
	fmt.Println("v:", v)
	fmt.Println("v Type:", v.Type())
	fmt.Println("v canSet:", v.CanSet())

	e := v.Index(0)
	fmt.Println("e Type:", e.Type())
	fmt.Println("e canSet:", e.CanSet())
	e.SetInt(999)
	fmt.Printf("%+v\n", a)
}
