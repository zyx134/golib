package _md5

import (
	"crypto/md5"
	"fmt"
)

//生成md5
func Md5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
