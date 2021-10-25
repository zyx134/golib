package example

import "testing"

func TestReadConfig(t *testing.T) {
	ReadConfig("./config.ini")
}
func TestWriteConfig(t *testing.T) {
	var conf Config
	conf.ServerConfig.Port = 8001
	conf.ServerConfig.Ip = "127.0.0._1"
	conf.MysqlConfig.Port = 3306
	WriteConfig("./config1.ini", conf)
}

func TestReadConfig2(t *testing.T) {
	conf := Config{
		ServerConfig{Port: 80, Ip: "127.0.0.2"},
		MysqlConfig{
			Username: "abc",
			Password: "1231asd",
			Database: "abc",
			Port:     3306,
			Host:     "127.0.01",
			Timeout:  100,
		},
	}
	WriteConfig("./config2.ini", conf)
}
