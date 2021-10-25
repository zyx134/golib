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
