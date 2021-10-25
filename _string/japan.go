package _string

import "regexp"

// 检查是否包含日语的正则表达式
var chkContainJPReg, _ = regexp.Compile(`[\x{0800}-\x{4e00}]+?`)

// 检查是否包含日语
func ChkContainJP(srcStr string) bool {
	if "" == srcStr {
		return false
	}

	if true == chkContainJPReg.MatchString(srcStr) {
		return true
	}

	return false
}
