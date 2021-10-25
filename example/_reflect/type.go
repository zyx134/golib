package _reflect

import (
	"fmt"
	"reflect"
)

//反射获取interface类型信息
func GetType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf(" %v 类型是：%v\n", a, t)
	// kind()可以获取具体类型
	k := t.Kind()
	switch k {
	case reflect.Int:
		fmt.Printf("%v is int\n", a)
	case reflect.Int64:
		fmt.Printf("%v is int64\n", a)
	case reflect.Float64:
		fmt.Printf("%v is float64\n", a)
	case reflect.String:
		fmt.Printf("%v is string\n", a)
	case reflect.Slice:
		fmt.Printf("%v is slice\n", a)
	case reflect.Bool:
		fmt.Printf("%v is bool\n", a)
	case reflect.Map:
		fmt.Printf("%v is map\n", a)
	case reflect.Func:
		fmt.Printf("%v is func\n", a)
	default:
		fmt.Printf("%v is undefind\n", a)
	}
}
