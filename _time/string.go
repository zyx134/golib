package _time

import "time"

var (
	TimeTemplate map[string]string = map[string]string{
		//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
		"timeTemplate1": "2006-01-02 15:04:05", //常规类型
		"timeTemplate2": "2006/01/02 15:04:05", //其他类型
		"timeTemplate3": "2006-01-02",          //其他类型
		"timeTemplate4": "15:04:05",            //其他类型
	}
)

//获取当前时间字符串
func String(model string) string {
	if model == "" {
		model = "2006-01-02 15:04:05"
	}
	timeStr := time.Now().Format(model)
	return timeStr
}
