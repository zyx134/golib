package _time

import "time"

//获取时间戳 int64
func Stamp64() int64 {
	return time.Now().Unix()
}

//正确的使用方式
//time.Sleep(time.Second * time.Duration(common.C.CpuTempRest))
//获取指定时间的时间戳
//获取一个月之后的时间戳
//30 * 24 * 60 * 60 * time.Second
//或者
//30 * 24 * time.Hour
func AppointStamp64(t time.Duration) (timestamp int64) {
	return time.Now().Add(t).Unix()
}

//将时间戳转换为字符串
func Int64ToString(t int64) string {
	return time.Unix(t, 0).Format(TimeTemplate["timeTemplate1"])
}
