package _string

import "strconv"

//字符串 to int32
//内部通过调用 ParseInt(s, 10, 0) 来实现的
func ToInt(from string) (int, error) {
	return strconv.Atoi(from)
}
