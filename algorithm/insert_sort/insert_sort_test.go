package inset_sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(s)
	InsertSort(s)
	fmt.Println(s)
}
func TestInsertSort1(t *testing.T) {
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(s)
	InsertSort1(s)
	fmt.Println(s)
}
