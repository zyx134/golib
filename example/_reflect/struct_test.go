package _reflect

import "testing"

func TestStruct(t *testing.T) {
	Struct(User{Id: 1, Name: "abc", Age: 99})
}

func TestGetStruct(t *testing.T) {
	GetStruct()
}
func TestSetStructValue(t *testing.T) {
	SetStructValue(&User{11, "abc", 88})
}

func TestCallStructMethod(t *testing.T) {
	CallStructMethod(User{11, "abc", 88}, "非指针")
	CallStructMethod(&User{11, "abc", 88}, "指针")
}
