package _reflect

import (
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	if obj1.Kind() != reflect.Struct {
		return data
	}
	for i := 0; i < obj1.NumField(); i++ {
		fieldValue := obj2.FieldByName(obj1.Field(i).Name)
		if fieldValue.CanInterface() {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
		//data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
