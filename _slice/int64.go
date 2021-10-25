package _slice

import "strconv"

// 过滤重复元素
func UniqueInt64(slc []int64) []int64 {
	strS := make([]string, len(slc))

	for k, v := range slc {
		strS[k] = strconv.FormatInt(v, 10)
	}

	strS = Unique(strS)

	slc = slc[0:0:len(strS)]

	for _, v := range strS {
		int64V, _ := strconv.ParseInt(v, 10, 64)

		slc = append(slc, int64V)
	}

	return slc
}

// 返回存在于 s1, 不存在于 s2 的值
// 返回值是一个新切片
func DiffInt64(s1 []int64, s2 []int64) []int64 {
	if 0 == len(s2) {
		return s1
	}

	s2Map := make(map[int64]int64)

	for _, v := range s2 {
		s2Map[v] = v
	}

	diff := make([]int64, 0, len(s1))
	for _, v := range s1 {
		if _, ok := s2Map[v]; ok {
			continue
		}

		diff = append(diff, v)
	}

	return diff
}
