package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 56, 12, 32, 4656, 2, 111, 1, 3345, 66, 221, 3456, 90, 11}
	fmt.Println(QuickSort(arr))
}
