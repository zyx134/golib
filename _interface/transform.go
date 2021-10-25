package _interface

import (
	"strconv"
)

// 解析int8数据
func Int8(o interface{}) int8 {
	return int8(Int64(o))
}

// 解析int数据
func Int(o interface{}) int {
	return int(Int64(o))
}

// 解析int数据
func Int32(o interface{}) int32 {
	return int32(Int64(o))
}

// 解析int64数据
func Int64(o interface{}) int64 {
	switch o_ := o.(type) {
	case int8:
		return int64(o_)
	case uint8:
		return int64(o_)
	case int:
		return int64(o_)
	case uint:
		return int64(o_)
	case int32:
		return int64(o_)
	case uint32:
		return int64(o_)
	case int64:
		return o_
	case uint64:
		return int64(o_)
	case string:
		i, _ := strconv.ParseInt(o_, 10, 64)
		return i
	case float32:
		return int64(o_)
	case float64:
		return int64(o_)
	default:
		return 0
	}
}

// 解析int64数据
func Uint64(o interface{}) uint64 {
	switch o_ := o.(type) {
	case int8:
		return uint64(o_)
	case uint8:
		return uint64(o_)
	case int:
		return uint64(o_)
	case uint:
		return uint64(o_)
	case int32:
		return uint64(o_)
	case uint32:
		return uint64(o_)
	case int64:
		return uint64(o_)
	case uint64:
		return o_
	case string:
		i, _ := strconv.ParseUint(o_, 10, 64)
		return i
	case float32:
		return uint64(o_)
	case float64:
		return uint64(o_)
	default:
		return 0
	}
}

// 解析float64数据
func Float64(o interface{}) float64 {
	switch o.(type) {
	case float64:
		return o.(float64)
	case string:
		if "" != o {
			res, _ := strconv.ParseFloat(o.(string), 10)
			return res
		}
		return 0
	case int:
		return float64(o.(int))
	case int64:
		return float64(o.(int64))
	default:
		return 0
	}
}

func String(o interface{}) string {
	switch o.(type) {
	case string:
		return o.(string)
	default:
		return ""
	case bool:
		if o.(bool) == false {
			return "false"
		}
		return "true"
	}
}

func IntSlice(o interface{}) []int {
	switch o.(type) {
	case []int:
		return o.([]int)
	default:
		return nil
	}
}

func Float64Slice(o interface{}) []float64 {
	switch o.(type) {
	case []float64:
		return o.([]float64)
	default:
		return nil
	}
}
