package redis

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

// hset(key, field, value)：向名称为key的hash中添加元素field
// hget(key, field)：返回名称为key的hash中field对应的value
// hmget(key, (fields))：返回名称为key的hash中field i对应的value
// hmset(key, (fields))：向名称为key的hash中添加元素field
// hexists(key, field)：名称为key的hash中是否存在键为field的域
// hincrby(key, field, integer)：将名称为key的hash中field的value增加integer
// hdel(key, field)：删除名称为key的hash中键为field的域
// hlen(key)：返回名称为key的hash中元素个数
// hkeys(key)：返回名称为key的hash中所有键
// hvals(key)：返回名称为key的hash中所有键对应的value
// hgetall(key)：返回名称为key的hash中所有的键（field）及其对应的value

// "test","e",1
func HashSet(r *redis.Pool, key, field string, value interface{}) error {
	conn := r.Get()
	defer conn.Close()
	_, err := conn.Do("hset", key, field, value)
	if err != nil {
		return err
	}
	return nil
}

// "test","a"
func HashGet(r *redis.Pool, key, field string) (b []byte, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err := conn.Do("hget", key, field)
	if err != nil {
		return
	}
	if res == nil {
		return b, errors.New("没有获取到数据")
	}
	b = res.([]byte)
	return
}

// "test","a"
func HashGetString(r *redis.Pool, key, field string) (string, error) {
	b, err := HashGet(r, key, field)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//values 是由 key, k,v,k,v组成
// "test", "a", "aa"
func HashMset(r *redis.Pool, values ...interface{}) error {
	conn := r.Get()
	defer conn.Close()
	_, err := conn.Do("hmset", values...)
	if err != nil {
		return err
	}
	return nil
}

//values 是由 key, k,k,k组成
//输入 "test", "a", "b","c"
//输出
func HashMget(r *redis.Pool, values ...interface{}) (b []interface{}, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err := redis.Values(conn.Do("hmget", values...))
	if err != nil {
		return
	}
	return res, nil
}

//values 是由 key, k,k,k组成
//输入 "test", "a", "b","c"
//输出 ["aa","bb","cc"]
func HashMgetString(r *redis.Pool, values ...interface{}) (b []string, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err := redis.Values(conn.Do("hmget", values...))
	if err != nil {
		return
	}
	for k := range res {
		b = append(b, string(res[k].([]byte)))
	}
	return
}

//values 是由 key, k,k,k组成
//输入 "test", "a", "b","c","cc","e"
//输出
// map[string]string{
// "a": "aa",
// "b": "bb",
// "c": "cc",
// "cc": "cc",
// "e": "3"
// }
func HashMgetMapString(r *redis.Pool, values ...interface{}) (b map[string]string, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err := redis.Values(conn.Do("hmget", values...))
	if err != nil {
		return
	}
	b = make(map[string]string)
	for k := range values {
		if k == 0 {
			continue
		}
		if val, ok := res[k-1].([]byte); ok {
			b[values[k].(string)] = string(val)
		}
	}
	return
}

//判断field是否存在
//输入 "test", "a"
//输出 true or false
func HashHexists(r *redis.Pool, key, field string) (ok bool) {
	conn := r.Get()
	defer conn.Close()
	res, err := conn.Do("hexists", key, field)
	if err != nil {
		return
	}
	if i := res.(int64); i == 1 {
		ok = true
	}
	return
}

//value 增加incr e原来是1，incr=2,完成后e=3
//要求 field的值必须是int
//若field不存在，则新增field,value=incr
//输入 "test", "e",2
//输出 true or false
func HashHincrby(r *redis.Pool, key, field string, incr int) (err error) {
	conn := r.Get()
	defer conn.Close()
	_, err = conn.Do("hincrby", key, field, incr)
	if err != nil {
		return
	}
	return
}

//删除field
//删除不存在的field不会报错
//输入 "test", "e"
func HashHdel(r *redis.Pool, key, field string) (err error) {
	conn := r.Get()
	defer conn.Close()
	_, err = conn.Do("hdel", key, field)
	if err != nil {
		return
	}
	return
}

//获取key的field长度
//输入 "test"
func HashHlen(r *redis.Pool, key string) (l int64, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err := conn.Do("hlen", key)
	if err != nil {
		return
	}
	l = res.(int64)
	return
}

//获取key的所有field
//输入 "test"
func HashHkeys(r *redis.Pool, key string) (l []string, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err := redis.Values(conn.Do("hkeys", key))
	if err != nil {
		return
	}
	for k := range res {
		if val, ok := res[k].([]byte); ok {
			l = append(l, string(val))
		}
	}
	return
}

//获取key的所有field的value
//输入 "test"
func HashHvals(r *redis.Pool, key string) (l []string, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err := redis.Values(conn.Do("hvals", key))
	if err != nil {
		return
	}
	for k := range res {
		if val, ok := res[k].([]byte); ok {
			l = append(l, string(val))
		}
	}
	return
}

//获取key的全部数据
//输入 "test"
func HashHgetallMapStringString(r *redis.Pool, key string) (res map[string]string, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err = redis.StringMap(conn.Do("hgetall", key))
	if err != nil {
		return
	}
	return
}

//获取key的全部数据
//输入 "test"
func HashHgetallMapStringInt(r *redis.Pool, key string) (res map[string]int, err error) {
	conn := r.Get()
	defer conn.Close()
	res, err = redis.IntMap(conn.Do("hgetall", key))
	if err != nil {
		return
	}
	return
}
