package _string

import (
	"github.com/Chain-Zhang/pinyin"
	"regexp"
	"strings"
	"unicode"
)

//中文转拼音
//Mode 介绍
// InitialsInCapitals: 首字母大写, 不带音调
// WithoutTone: 全小写,不带音调
// Tone: 全小写带音调
func ChineseToPinyin(str, split string) (ret string, err error) {
	ret, err = pinyin.New(str).Split(split).Mode(pinyin.InitialsInCapitals).Convert()
	return
}

//截取中文字符串
func SubChineseString(str string, length int) (ret string) {
	var retRune []rune
	for _, char := range str {
		if unicode.Is(unicode.Han, char) {
			retRune = append(retRune, char)
		}
	}
	if length > len(retRune) {
		length = len(retRune)
	}
	ret = string(retRune[0:length])
	return
}

// 移除字符串中的中文字符
func RemoveChinese(str string) string {
	str = strings.TrimSpace(str)
	if "" == str {
		return ""
	}

	strS := make([]string, 0, len(str))

	valid, _ := regexp.Compile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]")
	for _, r := range str {
		if true == unicode.Is(unicode.Scripts["Han"], r) {
			continue
		}

		if true == valid.MatchString(string(r)) {
			continue
		}

		strS = append(strS, string(r))
	}

	return strings.Join(strS, "")
}

// 检查字符串中是否包含中文字符
func HasChinese(str string) bool {
	if "" == str {
		return false
	}

	for _, r := range str {
		if true == unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}

// 检查字符串中是否包含非中文字符
func HasNonChinese(str string) bool {
	if "" == str {
		return true
	}

	for _, r := range str {
		if false == unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}
