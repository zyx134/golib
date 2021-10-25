package _byte

import (
	"bytes"
	"strconv"
)

func ToInt(b []byte) (i int, err error) {
	bytesStr := bytes.NewBuffer(b).String()
	i, err = strconv.Atoi(bytesStr)
	return
}
func ToString(b []byte) string {
	return string(b)
}
