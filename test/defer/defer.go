package _defer

import "fmt"

func def() {
	go func() {
		fmt.Println("前")
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
	}()
	defer fmt.Println("中")
	defer fmt.Println("后")
	panic("123")
}
