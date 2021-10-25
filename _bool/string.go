package _bool

import "strconv"

// 接受 _1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
// 其他形式的字符串会返回错误
func ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}
