package _string

import "strconv"

//字符串 to int64
func ToInt64(from string) (int64, error) {
	return strconv.ParseInt(from, 10, 64)
}

/**
参数 base 代表字符串按照给定的进制进行解释。一般的，base 的取值为 2~36，如果 base 的值为 0，则会根据字符串的前缀来确定 base 的值："0x" 表示 16 进制； "0" 表示 8 进制；否则就是 10 进制。

参数 bitSize 表示的是整数取值范围，或者说整数的具体类型。取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。

这里有必要说一下，当 bitSize==0 时的情况。

Go 中，int/uint 类型，不同系统能表示的范围是不一样的，目前的实现是，32 位系统占 4 个字节；64 位系统占 8 个字节。当 bitSize==0 时，应该表示 32 位还是 64 位呢？这里没有利用 runtime.GOARCH 之类的方式，而是巧妙的通过如下表达式确定 intSize：

const intSize = 32 << uint(^uint(0)>>63)
const IntSize = intSize // number of bits in int, uint (32 or 64)
主要是 ^uint(0)>>63 这个表达式。操作符 ^ 在这里是一元操作符 按位取反，而不是 按位异或。更多解释可以参考：Go 位运算：取反和异或。
*/
func ParseInt(s string, base int, bitSize int) (i int64, err error) {
	return strconv.ParseInt(s, base, bitSize)
}
