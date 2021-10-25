package _string

import (
	"bytes"
	"strings"
)

//高效字符串拼接
func StringJoin(strs ...string) (ret string) {
	if len(strs) == 0 {
		return
	}
	var buf bytes.Buffer
	for _, v := range strs {
		buf.WriteString(v)
	}
	return buf.String()
}

//查询字符串多条件包含判断
//存在其中一个则返回true
func ManyContains(str string, keys ...string) bool {
	for _, v := range keys {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

//截取字符串
//获取最后一份字符串
//例如 /a/b/c/de/f  获取f
func StrGetLastImpStr(str, exp string) string {
	index := strings.LastIndex(str, exp)
	return str[index+1:]
}

// 字符串转换为大写, 不需要转换的字符或者字符串不转化
func ToUpperSpecialAndNoNeed(str *string, noCaseStr string) string {
	if false == strings.Contains(*str, noCaseStr) {
		*str = strings.ToUpper(*str)

		return *str
	}

	strS := strings.Split(*str, noCaseStr)

	for k, v := range strS {
		v = strings.ToUpper(v)

		strS[k] = v
	}

	*str = strings.Join(strS, noCaseStr)

	return *str
}

// 返回字符串，该字符串为了数据库查询语句等的需要在某些字符前加上了反斜线。
// 这些字符是单引号（'）、双引号（"）、反斜线（\）与 NUL（NULL 字符）。
func AddSlashes(str string) string {
	if "" == str {
		return str
	}

	// 拆分成单个字符的切片
	s := strings.Split(str, "")

	// 判断字符串是否 ' " \, 是的话加上转义
	for k, char := range s {
		// 字符串是 ' " \, 加上转义
		if "'" == char || "\"" == char || "\\" == char {
			char = "\\" + char
			s[k] = char

			continue
		}

		// 字符串是 \x00(空字符串), 重置为空字符串
		if "\x00" == char {
			char = ""
		}

		s[k] = char
	}

	return strings.Join(s, "")
}

// ContactStr 字符串的拼接
func ContactStr(str ...string) string {
	var buffer bytes.Buffer

	for _, s := range str {
		buffer.WriteString(s)
	}
	return buffer.String()
}

// Substr returns the substr from start to length.
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// Addslashes
func Addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

//SubstrBetween
// ./abc/asd.txt
//我需要 abc
//截取两个字符串中间的字符串
func SubstrBetween(str, s1, s2 string) string {
	c := str
	index := strings.Index(str, s1)
	if index == -1 {
		return c
	}
	str = str[index+len(s1):]
	index = strings.Index(str, s2)
	if index == -1 {
		return c
	}
	return str[:index]
}
