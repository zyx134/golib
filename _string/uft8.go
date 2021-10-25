package _string

import "unicode/utf8"

//判断是不是utf8，如果不是就转换成utf8
//接收字节数组，返回字符串
func IsUtf8(b []byte) string {
	var strs string
	isUtf8 := utf8.Valid(b)
	if !isUtf8 {
		strs = Convert(string(b), "gbk", "utf8")
	} else {
		strs = string(b)
	}
	return strs
}
