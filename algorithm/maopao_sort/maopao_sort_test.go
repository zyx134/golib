package maopao_sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println(BubbleSort([]int{22, 4565, 123, 11, 35, 5, 123, 89, 12234, 788, 123123, 1, 0, 999, 22}))
}

func TestBubbleSortGetMax(t *testing.T) {
	fmt.Println("最大值是：", BubbleSortGetMax([]int{22, 4565, 123, 11, 35, 5555555, 123, 89, 12234, 788, 123123, 1, 0, 999, 22}))
}
