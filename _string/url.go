package _string

import (
	"net/url"
	"strings"
)

// Urlencode - Encodes string
// 和 php 的 urlencode 函数保持一致
func Urlencode(s string) string {
	s = url.QueryEscape(s)
	s = strings.Replace(s, "~", "%7E", -1)

	return s
}

// Urldecode - Decodes URL-encoded string
// 和 php 的 urldecode 函数保持一致
func Urldecode(s string) (string, error) {
	s = strings.Replace(s, "%7E", "~", -1)

	return url.QueryUnescape(s)
}

// Rawurlencode - Raw-URL-encode string
func RawUrlencode(s string) string {
	s = url.QueryEscape(s)

	return strings.Replace(s, "+", "%20", -1)
}

// Rawurldecode - Decodes Raw-URL-encoded string
func RawUrldecode(s string) (string, error) {
	strings.Replace(s, "%20", "+", -1)

	return url.QueryUnescape(s)
}
