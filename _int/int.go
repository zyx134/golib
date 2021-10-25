package _int

import (
	"math/rand"
	"time"
)

///获取随机数
///开始 结束  取的数量
//生成[15，88]之间的随机数,括号左包含右不包含
//现在是右边也包含
//n:=rand.Intn(73)+15 //(88-15 )+15
func IntRandom(start int, end int, num int) (rands []int) {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	for i := 0; i < num; i++ {
		rands = append(rands, rand.Intn(end-start+1)+start)
	}
	return
}

//int32 to int64
func IntToInt64(from int) int64 {
	return int64(from)
}
