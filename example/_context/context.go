package _context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//基本案例
func Context() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go worker(ctx, wg)
	time.Sleep(time.Second * 3)
	wg.Wait()
	fmt.Println("over")
}
func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:

		}
	}
}
func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}
