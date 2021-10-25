package example

type Config struct {
	ServerConfig ServerConfig `ini:"server"`
	MysqlConfig  MysqlConfig  `ini:"mysql"`
}
type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}
type MysqlConfig struct {
	Username string  `ini:"username"`
	Password string  `ini:"password"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}
