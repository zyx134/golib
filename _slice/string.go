package _slice

import "strings"

// 滤重复元素
func Unique(slc []string) []string {
	// 切片长度小于1024的时候，循环来过滤
	if len(slc) < 1024 {
		return removeRepByLoop(slc)
	}

	// 大于的时候，通过map来过滤
	return removeRepByMap(slc)
}

// 通过两重循环过滤重复元素
func removeRepByLoop(slc []string) []string {
	// 存放结果
	var result []string

	for i := range slc {
		flag := true
		for j := range result {
			// 存在重复元素，标识为false
			if slc[i] == result[j] {
				flag = false
				break
			}
		}

		// 标识为false，不添加进结果
		if flag {
			result = append(result, slc[i])
		}
	}

	return result
}

// 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []string) []string {
	var result []string

	// 存放不重复主键
	tempMap := map[string]byte{}

	for _, e := range slc {
		l := len(tempMap)

		tempMap[e] = 0

		// 加入map后，map长度变化，则元素不重复
		if len(tempMap) != l {
			result = append(result, e)
		}
	}

	return result
}

// 检查字符串是否在切片中
func InArrayStr(needle *string, haystack []string) bool {
	if 0 == len(haystack) {
		return false
	}

	for _, v := range haystack {
		if *needle == v {
			return true
		}
	}

	return false
}

//插入指定位置
func InsertToSliceString(s []string, index int) []string {
	temp := append([]string{}, s[index:]...)
	s = append(s[:index], "ee")
	s = append(s, temp...)
	return s //[aa bb ee cc dd]
}

//删除指定索引
func DelIndexSliceString(s []string, index int) []string {
	s = append(s[:index], s[index+1:]...)
	return s //[aa bb cc dd]
}

//删除指定值
func DelValueSliceString(s []string, value string) []string {
	for k, v := range s {
		if v == value {
			s = append(s[:k], s[k+1:]...)
		}
	}
	return s //[aa bb cc]
}

// 返回存在于 s1, 不存在于 s2 的值
// 返回值是一个新切片
// 不区分大小写
func DiffStr(s1 []string, s2 []string) []string {
	if 0 == len(s2) {
		return s1
	}

	s2Map := make(map[string]string)

	for _, v := range s2 {
		v = strings.ToUpper(v)
		s2Map[v] = v
	}

	diff := make([]string, 0, len(s1))
	for _, v := range s1 {
		v = strings.ToUpper(v)
		if _, ok := s2Map[v]; ok {
			continue
		}

		diff = append(diff, v)
	}

	return diff
}

// 滤重复元素
func UniqueString(slc []string) []string {
	// 切片长度小于1024的时候，循环来过滤
	if len(slc) < 1024 {
		return removeRepByLoop(slc)
	}

	// 大于的时候，通过map来过滤
	return removeRepByMap(slc)
}
