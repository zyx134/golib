package _int

import "strconv"

//int32 to string
//Itoa 内部直接调用 FormatInt(i, 10)
func IntToString(from int) string {
	return strconv.Itoa(from)
}

//int64 to string
func Int64ToString(from int64) string {
	return strconv.FormatInt(from, 10)
}

// uint64转字符串
func FormatUint(i uint64, base int) string {
	return strconv.FormatUint(i, base)
}

// int64转字符串
func FormatInt(i int64, base int) string {
	return strconv.FormatInt(i, base)
}
