package _float

import (
	"math"
	"strconv"
)

//float64 to string
func Float64ToString(from float64) string {
	str := strconv.FormatFloat(from, 'E', -1, 64)
	return str
}

//fmt 就是fmt.Printf的占位符
//prec 表示有效数字（对 fmt='b' 无效），对于 'e', 'E' 和 'f'，有效数字用于小数点之后的位数；对于 'g' 和 'G'，则是所有的有效数字
//fmt='b' 时，得到的字符串是无法通过 ParseFloat 还原的。

//%b        无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat
//        的 'b' 转换格式一致。例如 -123456p-78
//%e        科学计数法，例如 -1234.456e+78                                    Printf("%e", 10.2)                            _1.020000e+01
//%E        科学计数法，例如 -1234.456E+78                                    Printf("%e", 10.2)                            _1.020000E+01
//%f        有小数点而无指数，例如 123.456                                    Printf("%f", 10.2)                            10.200000
//%g        根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出                Printf("%g", 10.20)                            10.2
//%G        根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出                Printf("%G", 10.20+2i)

func FormatFloat(f float64, fmt byte, prec, bitSize int) string {
	return strconv.FormatFloat(f, fmt, prec, bitSize)
}

// 金额相除
func MoneyDivide(left, right float64, round int) float64 {
	if 0 == right {
		return 0
	}

	roundBit := 6

	left = left * math.Pow10(roundBit)

	res := Round(left/right, 0)

	res = Round(res/math.Pow10(roundBit), round)

	return res
}

// 金额相乘
func MoneyMultiply(left, right float64, round int) float64 {
	roundBit := 6

	left = left * math.Pow10(roundBit)

	res := Round(left*right, 0)

	res = Round(res/math.Pow10(roundBit), round)

	return res
}

// 金额相加
func MoneyAdd(left, right float64, round int) float64 {
	roundBit := 6

	left = left * math.Pow10(roundBit)
	right = right * math.Pow10(roundBit)

	res := left + right

	res = Round(res/math.Pow10(roundBit), round)

	return res
}

// 金额相减
func MoneySubtract(left, right float64, round int) float64 {
	roundBit := 6

	left = left * math.Pow10(roundBit)
	right = left * math.Pow10(roundBit)

	res := left - right

	res = Round(res/math.Pow10(roundBit), round)

	return res
}

// 四舍五入
// todo 增加 round mode
func Round(x float64, round int) float64 {
	r := math.Pow10(round)

	x = math.Floor((x*r)+0.50000000001) / r

	return x
}

// 金额舍去 round 位之后的数字
func MoneyFloor(x float64, round int) float64 {
	l := math.Pow10(round)

	x = math.Floor(x*l) / l

	return x
}
