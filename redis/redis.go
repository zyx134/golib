package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/zyx134/golib/_string"

	"github.com/gomodule/redigo/redis"
)

type Config struct {
	Host      string
	Password  string
	Db        string
	Port      int
	MaxIdle   int
	MaxActive int
	Timeout   time.Duration
}

func New(c Config) *redis.Pool {
	r := new(redis.Pool)
	r.MaxIdle = c.MaxIdle
	r.MaxActive = c.MaxActive
	r.IdleTimeout = c.Timeout
	r.Dial = func() (redis.Conn, error) {
		cc, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
		if err != nil {
			return nil, err
		}
		if c.Password != "" {
			if _, err := cc.Do("AUTH", c.Password); err != nil {
				cc.Close()
				return nil, err
			}
		}
		_, err = cc.Do("SELECT", c.Db)
		if err != nil {
			cc.Close()
			return nil, err
		}
		return cc, err
	}
	r.TestOnBorrow = func(c redis.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	}

	return r
}
func Exists(r *redis.Pool, key string) bool {
	conn := r.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Delete(r *redis.Pool, key string) (bool, error) {
	conn := r.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(r *redis.Pool, key string) error {
	conn := r.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(r, key)
		if err != nil {
			return err
		}
	}

	return nil
}
func Set(r *redis.Pool, key string, data interface{}, time int) (bool, error) {
	conn := r.Get()
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

func Get(r *redis.Pool, key string) (string, error) {
	conn := r.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
func SetInt(r *redis.Pool, key string, data int, time int) (bool, error) {
	conn := r.Get()
	defer conn.Close()
	reply, err := redis.Bool(conn.Do("SET", key, data))
	conn.Do("EXPIRE", key, time)

	return reply, err
}
func GetInt(r *redis.Pool, key string) (ret int, err error) {
	s, err := Get(r, key)
	if err != nil {
		return
	}
	return _string.ToInt(s)
}
