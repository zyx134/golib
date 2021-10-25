package _string

import (
	"time"

	"github.com/zyx134/golib/_int"
	"github.com/zyx134/golib/_time"
)

//字符串转时间
func ToTime(str, model string) (time.Time, error) {
	//使用parseInLocation将字符串格式化返回本地时区时间
	time, err := time.ParseInLocation(_time.TimeTemplate[model], str, time.Local)
	return time, err
}

//字符串转时间戳
func ToTimeStamp(str, model string) (int, error) {
	time, err := ToTime(str, model)
	if err != nil {
		return 0, err
	}
	ret, err := _int.Int64ToInt(time.Unix())
	if err != nil {
		return 0, err
	}
	return ret, err
}
