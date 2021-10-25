package _reflect

import (
	"fmt"
	"reflect"
)

//反射获取interface值信息
func GetValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Printf("%v 的值是 %v \n", a, v)
	k := v.Kind()
	fmt.Printf("%v 的类型是 %v \n", a, k)
}
