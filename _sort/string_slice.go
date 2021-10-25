package _sort

import "sort"

//文件排序
// 1.txt,100.txt,101.txt,2.txt
func FileSort(filelist []string) []string {
	sort.Slice(
		filelist,
		func(i, j int) bool {
			return SortName(filelist[i]) < SortName(filelist[j])
		},
	)
	return filelist
}
