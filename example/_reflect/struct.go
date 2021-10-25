package _reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(arg string) {
	fmt.Println("Hello ", arg)
}

//基础结构体
func Struct(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	//获取值
	v := reflect.ValueOf(a)
	fmt.Printf("值：%+v\n", v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		//获取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v\n", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

//查看匿名字段
func GetStruct() {
	m := Boy{User{1, "bob", 17}, "bg"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

//修改结构体的值
//取结构体地址 v := reflect.ValueOf(&a)
//获取结构体的具体值:Elem()
//获取结构体的字段:FieldByName()
func SetStructValue(a interface{}) {
	v := reflect.ValueOf(a)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("jack")
		fmt.Println("修改值成功")
	}
	fmt.Printf("%+v\n", a)
}

//调用方法
//指针类型和非指针类型都可以调用
func CallStructMethod(a interface{}, arg string) {
	v := reflect.ValueOf(a)
	// 获取方法
	m := v.MethodByName("Hello")
	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf(arg)}
	m.Call(args)
}
