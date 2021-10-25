package _context

import (
	"context"
	"fmt"
	"time"
)

//返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
//
//取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
func WithDeadline(d time.Time) {
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	select {
	case <-time.After(time.Second):
		fmt.Println("over")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
