package _context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

//WithValue返回父节点的副本，其中与key关联的值为val。
//
//仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。
//
//所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。WithValue的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。或者，导出的上下文关键变量的静态类型应该是指针或接口。
func WithValue() {
	wg := &sync.WaitGroup{}
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*50)
	defer cancel()
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "300")
	wg.Add(1)
	go worker3(ctx, wg)
	time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Println("over")
}

//类型转换，可以用断言(只能使用在interface{}类型转换成其它类型)
//显式类型强制转换（常规是用于基本类型）
func worker3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker done!")
}
