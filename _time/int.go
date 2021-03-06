package _time

import "github.com/zyx134/golib/_int"

//获取时间戳 int32
func Stamp32() (timestamp int) {
	t64 := Stamp64()
	timestamp, err := _int.Int64ToInt(t64)
	if err != nil {
		return 0
	}
	return timestamp
}
