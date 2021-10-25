package _reflect

import "reflect"

//判断两个变量是否相等
func CheckEq(x interface{}, y interface{}) bool {
	return reflect.DeepEqual(x, y)
}
