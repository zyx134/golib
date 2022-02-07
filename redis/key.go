package redis

import "github.com/gomodule/redigo/redis"

func (p *Pool) KeyExists(key string) bool {
	conn := p.pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func (p *Pool) KeyDelete(key string) (bool, error) {
	conn := p.pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
func (p *Pool) KeyLikeDeletes(key string) error {
	conn := p.pool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = p.KeyDelete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
