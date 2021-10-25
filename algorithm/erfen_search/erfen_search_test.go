package erfen_search

import (
	"fmt"
	"testing"
)

func TestBinSearch(t *testing.T) {
	arr := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i + 1
	}
	index := BinSearch(arr, 99)
	fmt.Println(arr)
	if index != -1 {
		fmt.Println(index, arr[index])
	} else {
		fmt.Println("没有找到数据")
	}
}
