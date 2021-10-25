package _string

import "strings"

//explode
//字符串打散成字符串切片
func Explode(str, split string) []string {
	strSlice := strings.Split(str, split)
	return strSlice
}

//implode
//字符串切片组合成字符串
func Implode(strSlice []string, split string) string {
	str := strings.Join(strSlice, split)
	return str
}
