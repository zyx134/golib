package inset_sort

//升序
func InsertSort(s []int) {
	l := len(s)
	if l < 2 {
		return
	}
	for i := 1; i < l; i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			swap(s, j, j-1)
		}
	}
}

//降序
func InsertSort1(s []int) {
	l := len(s)
	if l < 2 {
		return
	}
	for i := 1; i < l; i++ {
		for j := i; j > 0 && s[j] > s[j-1]; j-- {
			swap(s, j, j-1)
		}
	}
}
func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
