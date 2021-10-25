package _string

import "strconv"

//字符串 to float64
func ToFloat64(from string) (float64, error) {
	f64, err := strconv.ParseFloat(from, 64)
	return f64, err
}
