package _os

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
)

//获取文件列表 单层 不递归
func GetOneFileList(path string) (list []string, err error) {
	filelist, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for i := range filelist {
		list = append(list, filelist[i].Name())
	}
	return
}

//获取文件列表 单层 并且排序
func GetOneFileListAndSort(path string) (list []string, err error) {
	filelist, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for i := range filelist {
		list = append(list, filelist[i].Name())
	}
	sort.Slice(
		filelist,
		func(i, j int) bool {
			return sortName(filelist[i].Name()) < sortName(filelist[j].Name())
		},
	)
	return
}

// sortName returns a filename sort key with
// non-negative integer suffixes in numeric order.
// For example, amt, amt0, amt2, amt10, amt099, amt100, ...
func sortName(filename string) string {
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	// split numeric suffix
	i := len(name) - 1
	for ; i >= 0; i-- {
		if '0' > name[i] || name[i] > '9' {
			break
		}
	}
	i++
	// string numeric suffix to uint64 bytes
	// empty string is zero, so integers are plus one
	b64 := make([]byte, 64/8)
	s64 := name[i:]
	if len(s64) > 0 {
		u64, err := strconv.ParseUint(s64, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(b64, u64+1)
		}
	}
	// prefix + numeric-suffix + ext
	return name[:i] + string(b64) + ext
}

//获取文件列表 递归 获取完整路径
func GetAllFileList(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read {
		if fi.IsDir() { // 判断是否是文件夹
			fullDir := path + "/" + fi.Name() //构造新的路径
			// files = append(files, fullDir)            //追加路径
			files, _ = GetAllFileList(fullDir, files) //文件夹递归
		} else {
			fullDir := path + "/" + fi.Name() //构造新的路径
			files = append(files, fullDir)    //追加路劲
		}
	}
	return files, nil
}

//获取文件列表 递归 获取完整路径
//并且排序
func GetAllFileListAndSort(path string, files []string) ([]string, error) {
	filelist, err := GetAllFileList(path, files)
	if err != nil {
		return filelist, err
	}
	sort.Slice(
		filelist,
		func(i, j int) bool {
			return sortName(filelist[i]) < sortName(filelist[j])
		},
	)
	return filelist, nil
}

//获取文件列表 递归 获取当前文件路径
//appendPath中间需要拼接的路径
//例如 完整路径是 /a/b/c.txt
//但是我不需要完整路径 我需要 /b/c.txt
//这里传入 "/b"
func GetAllFileListNotFullPath(path, appendPath string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read {
		if fi.IsDir() { // 判断是否是文件夹
			fullDir := path + "/" + fi.Name() //构造新的路径
			app := appendPath + "/" + fi.Name()
			files, _ = GetAllFileListNotFullPath(fullDir, app, files) //文件夹递归
		} else {
			app := appendPath + "/" + fi.Name()
			files = append(files, app)
		}
	}
	return files, nil
}

//获取文件列表 递归 获取当前文件路径
//并且排序
func GetAllFileListNotFullPathAndSort(path, appendPath string, files []string) ([]string, error) {
	filelist, err := GetAllFileListNotFullPath(path, appendPath, files)
	if err != nil {
		return filelist, err
	}
	sort.Slice(
		filelist,
		func(i, j int) bool {
			return sortName(filelist[i]) < sortName(filelist[j])
		},
	)
	return filelist, nil
}
