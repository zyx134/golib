package redis

import (
	"fmt"
	"time"

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
type Pool struct {
	pool *redis.Pool
}

func New(c Config) *Pool {
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

	return &Pool{
		pool: r,
	}
}
