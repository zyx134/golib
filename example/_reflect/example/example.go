package example

import (
	"fmt"
	"io/ioutil"
)

//反射练习
//任务：解析如下配置文件
//序列化：将结构体序列化为配置文件数据并保存到硬盘
//反序列化：将配置文件内容反序列化到程序的结构体
//配置文件有server和mysql相关配置

//读取配置文件
func ReadConfig(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("反序列化成功  conf: %#v\n  port: %#v\n", conf, conf.ServerConfig.Port)
}

//写入配置文件
func WriteConfig(filename string, conf Config) {
	err := MarshalFile(filename, conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("写入配置文件成功")
}
