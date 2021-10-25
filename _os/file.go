package _os

import (
	"io/ioutil"
	"os"
)

//获取文件信息
func FileInfo(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

//判断文件或者文件夹是否存在
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

//移动文件
func MoveFile(from string, to string) (err error) {
	err = os.Rename(from, to)
	return
}

//读取文件
func ReadFile(path string) (ret []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return ret, err
	}
	defer file.Close()
	ret, err = ioutil.ReadAll(file)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

//文件不存在创建文件，存在就打开
func CreateOrOpenFile(filePath string) (f *os.File, err error) {
	if !FileIsExisted(filePath) {
		f, err = os.Create(filePath)
	} else {
		f, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	}
	return f, err
}

//OpenFile 既能打开一个已经存在的文件，也能创建并打开一个新文件。
//OpenFile 是一个更一般性的文件打开函数，大多数调用者都应用 Open 或 Create 代替本函数
/**
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
const (
    // 单字符是被 String 方法用于格式化的属性缩写。
    ModeDir        FileMode = _1 << (32 - _1 - iota) // d: 目录
    ModeAppend                                     // a: 只能写入，且只能写入到末尾
    ModeExclusive                                  // l: 用于执行
    ModeTemporary                                  // T: 临时文件（非备份文件）
    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    ModeDevice                                     // D: 设备
    ModeNamedPipe                                  // p: 命名管道（FIFO）
    ModeSocket                                     // S: Unix 域 socket
    ModeSetuid                                     // u: 表示文件具有其创建者用户 id 权限
    ModeSetgid                                     // g: 表示文件具有其创建者组 id 的权限
    ModeCharDevice                                 // c: 字符设备，需已设置 ModeDevice
    ModeSticky                                     // t: 只有 root/ 创建者能删除 / 移动文件

    // 覆盖所有类型位（用于通过 & 获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有 Unix 权限位（用于通过 & 获取类型位）
)
*/
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}
