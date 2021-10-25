package _string

import "strconv"

func ToUint64(s string) (n uint64, err error) {
	return strconv.ParseUint(s, 10, 64)
}
func ParseUint(s string, base int, bitSize int) (n uint64, err error) {
	return strconv.ParseUint(s, base, bitSize)
}
