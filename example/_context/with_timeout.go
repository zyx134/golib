package _context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//WithTimeout返回WithDeadline(parent, time.Now().Add(timeout))。
//
//取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。具体示例如下：
func WithTimeout() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*50)
	defer cancel()
	wg.Add(1)
	go worker1(ctx, wg)
	time.Sleep(time.Second * 3)
	fmt.Println("over")
	wg.Wait()
}
func worker1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("db connecting ...")  // 假设正常连接数据库耗时10毫秒
		time.Sleep(time.Millisecond * 20) // 50毫秒后自动调用
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker done")
}
