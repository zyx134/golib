package _slice

import "strconv"

// 过滤重复元素
func UniqueInt8(slc []int8) []int8 {
	strS := make([]string, len(slc))

	for k, v := range slc {
		strS[k] = strconv.FormatInt(int64(v), 10)
	}

	strS = Unique(strS)

	slc = slc[0:0:len(strS)]

	for _, v := range strS {
		int64V, _ := strconv.ParseInt(v, 10, 64)

		slc = append(slc, int8(int64V))
	}

	return slc
}
