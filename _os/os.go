package _os

import (
	"bufio"
	"os"
)

//写入string到文件
func WriteToFile(f *os.File, str string) error {
	write := bufio.NewWriter(f)
	_, err := write.WriteString(str)
	if err != nil {
		return err
	}
	return write.Flush()
}

//覆盖文件内容
//先清空，再写入
//清空文件然后写入
func ClearUpAndWrite(filepath string, str string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	err = WriteToFile(file, str)
	if err != nil {
		return err
	}
	return nil
}
