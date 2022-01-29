package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/zyx134/golib/_string"

	"github.com/gomodule/redigo/redis"
)

var (
	RedisConn *redis.Pool
)

func Init(viper *viper.Viper) error {
	host := viper.GetString("redis.host")
	port := viper.GetInt("redis.port")
	RedisConn = &redis.Pool{
		MaxIdle:     viper.GetInt("redis.maxIdle"),
		MaxActive:   viper.GetInt("redis.maxActive"),
		IdleTimeout: viper.GetDuration("redis.idleTimeout"),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return nil, err
			}
			if viper.GetString("redis.password") != "" {
				if _, err := c.Do("AUTH", viper.GetString("redis.password")); err != nil {
					c.Close()
					return nil, err
				}
			}
			_, err = c.Do("SELECT", viper.GetString("redis.db"))
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}
func Exists(key string) bool {
	conn := RedisConn.Get()
	// defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	// defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	// defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
func Set(key string, data interface{}, time int) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return false, err
	}
	conn.Do("EXPIRE", key, time)

	return true, err
}

func Get(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
func SetInt(key string, data int, time int) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bool(conn.Do("SET", key, data))
	conn.Do("EXPIRE", key, time)

	return reply, err
}
func GetInt(key string) (ret int, err error) {
	s, err := Get(key)
	if err != nil {
		return
	}
	return _string.ToInt(s)
}
