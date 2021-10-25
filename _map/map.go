package _map

import (
	"github.com/zyx134/golib/_reflect"
)

//map比较
//其中方法一用到了反射，效率相对比较低，相差大约10倍。
func CheckEq(m1, m2 map[interface{}]interface{}, t int) bool {
	if t == 1 {
		return _reflect.CheckEq(m1, m2)
	} else {
		if len(m1) == len(m2) {
			for k1, v1 := range m1 {
				if v2, ok := m2[k1]; ok {
					if v1 != v2 {
						return false
					}
				} else {
					return false
				}
			}
			return true
		}
		return false
	}
}
