package _reflect

import (
	"fmt"
	"reflect"
)

//反射修改值
//定义一个整型变量，得到反射值，检查是否可以修改。
//必须是指针类型才能修改
//接着，通过反射获取整型变量地址的反射值，再通过Elem()得到指针指向的具体对象，这样就可以修改值
func SetValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Int:
		t := v.Type()
		fmt.Println("v Type:", t)
		fmt.Println("v CanSet:", v.CanSet())
	case reflect.Ptr:
		t := v.Type()
		fmt.Println("v Type:", t)
		fmt.Println("v CanSet:", v.CanSet())
		fmt.Println("v.Elem() CanSet:", v.Elem().CanSet())
		v.Elem().SetFloat(7.9)
		fmt.Println(a, "设置为", v.Elem().Float())
	default:
		fmt.Println(a, "设置失败")
	}
}
