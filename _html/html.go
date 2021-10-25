package _html

import (
	"regexp"
	"strings"
)

var regLower *regexp.Regexp
var regStyle *regexp.Regexp
var regScript *regexp.Regexp
var regHtml *regexp.Regexp
var regCRLF *regexp.Regexp

func init() {
	regLower, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	regStyle, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	regScript, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	regHtml, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	regCRLF, _ = regexp.Compile("\\s{2,}")
}

// 过滤html标签
func TrimHtml(src string) string {
	// 将HTML标签全转换成小写
	src = regLower.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除STYLE
	src = regStyle.ReplaceAllString(src, "")

	// 去除SCRIPT
	src = regScript.ReplaceAllString(src, "")

	// 去除所有尖括号内的HTML代码，并换成换行符
	src = regHtml.ReplaceAllString(src, "\n")

	// 去除连续的换行符
	src = regCRLF.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}
