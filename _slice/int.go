package _slice

import "strconv"

// Find获取一个切片并在其中查找元素。如果找到它，它将返回它的key，否则它将返回-1和一个错误的bool。
func FindInt(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// 过滤重复元素
func UniqueInt(slc []int) []int {
	strS := make([]string, len(slc))

	for k, v := range slc {
		strS[k] = strconv.FormatInt(int64(v), 10)
	}

	strS = Unique(strS)

	slc = slc[0:0:len(strS)]

	for _, v := range strS {
		int64V, _ := strconv.ParseInt(v, 10, 64)

		slc = append(slc, int(int64V))
	}

	return slc
}

// 返回存在于 s1, 不存在于 s2 的值
// 返回值是一个新切片
func DiffInt(s1 []int, s2 []int) []int {
	if 0 == len(s2) {
		return s1
	}

	s2Map := make(map[int]int)

	for _, v := range s2 {
		s2Map[v] = v
	}

	diff := make([]int, 0, len(s1))
	for _, v := range s1 {
		if _, ok := s2Map[v]; ok {
			continue
		}

		diff = append(diff, v)
	}

	return diff
}

// 获取最小值
func MinInt(s []int) int {
	min := 0
	if 0 == len(s) {
		return min
	}

	min = s[0]

	for _, v := range s {
		if min > v {
			min = v
		}
	}

	return min
}
