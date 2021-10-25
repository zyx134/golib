package maopao_sort

import "fmt"

//冒泡排序

//算法描述：冒泡算法，数组中前一个元素和后一个元素进行比较如果大于或者小于 前者就进行交换，最终返回最大或者最小都冒到数组的最后序列时间复杂度为 O(n^2).
//
//算法步骤
//从数组开头选择相邻两个元素进行比较，并进行交换
//不停向后移动
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//冒泡排序获取最大值
func BubbleSortGetMax(arr []int) int {
	//一次循环能保证最大的往右边移动
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
		}
	}
	fmt.Println(arr)
	return arr[len(arr)-1]
}
