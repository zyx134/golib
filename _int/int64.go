package _int

import (
	"fmt"
	"strconv"
)

var HexMap = map[int64]int64{
	10: 65,
	11: 66,
	12: 67,
	13: 68,
	14: 69,
	15: 70,
}

//int64 转int
func Int64ToInt(i int64) (int, error) {
	strInt64 := strconv.FormatInt(i, 10)
	return strconv.Atoi(strInt64)
}

// 10进制数字转16进制字符串
func Oct2Hex(n int64) string {
	if n < 0 {
		return ""
	}

	if n == 0 {
		return "0"
	}

	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = HexMap[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}

	return s
}
