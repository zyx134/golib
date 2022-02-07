package redis

import (
	"github.com/gomodule/redigo/redis"
)

//设置
func (p *Pool) SetSet(key, value string) error {
	conn := p.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SADD", key, value)
	if err != nil {
		return err
	}
	return nil
}

//设置
func (p *Pool) SetSetAndExpire(key, value string) error {
	conn := p.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SADD", key, value)
	if err != nil {
		return err
	}
	return nil
}

//判断是否存在
func (p *Pool) SetExist(key, value string) bool {
	conn := p.pool.Get()
	defer conn.Close()
	exist, err := redis.Int(conn.Do("SISMEMBER", key, value))
	if err != nil {
		return false
	}
	if exist == 1 {
		return true
	}
	return false
}

//获取全部数据
func (p *Pool) SetGetAll(key string) []string {
	conn := p.pool.Get()
	defer conn.Close()
	res, _ := redis.Strings(conn.Do("SMEMBERS", key))
	return res
}

//删除
func (p *Pool) SetDel(key, value string) error {
	conn := p.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SREM", key, value)
	return err
}

//随机删除
func (p *Pool) SetPop(key string, count int) []string {
	conn := p.pool.Get()
	defer conn.Close()
	res, _ := redis.Strings(conn.Do("SPOP", key, count))
	return res
}
