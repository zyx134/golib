package redis

import "github.com/gomodule/redigo/redis"

const (
	SET_IF_NOT_EXIST      = "NX" // 不存在则执行
	SET_WITH_EXPIRE_TIME  = "EX" // 过期时间(秒)  PX 毫秒
	SET_LOCK_SUCCESS      = "OK" // 操作成功
	DEL_LOCK_SUCCESS      = 1    // lock 删除成功
	DEL_LOCK_NON_EXISTENT = 0    // 删除lock key时,并不存在
)

/*
   redis 类型 字符串设置一个分布式锁 (哈希内部字段不支持过期判断,redis只支持顶级key过期)

   @param key: 锁名,格式为  用户id_操作_方法
   @param requestId:  客户端唯一id 用来指定锁不被其他线程(协程)删除
   @param ex: 过期时间
*/
func (p *Pool) LockAdd(key, requestId string, ex int) bool {
	conn := p.pool.Get()
	defer conn.Close()
	msg, _ := redis.String(
		conn.Do("SET", key, requestId, SET_IF_NOT_EXIST, SET_WITH_EXPIRE_TIME, ex),
	)
	return msg == SET_LOCK_SUCCESS
}

/*
   获得redis分布式锁的值

   @param key:redis类型字符串的key值
   @param return: redis类型字符串的value
*/
func (p *Pool) LockGet(key string) string {
	conn := p.pool.Get()
	defer conn.Close()
	msg, _ := redis.String(conn.Do("GET", key))
	return msg
}

/*
   删除redis分布式锁

   @param key:redis类型字符串的key值
   @param requestId: 唯一值id,与value值对比,避免在分布式下其他实例删除该锁
*/
func (p *Pool) LockDel(key, requestId string) bool {
	conn := p.pool.Get()
	defer conn.Close()
	if p.LockGet(key) == requestId {
		msg, _ := redis.Int64(conn.Do("DEL", key))
		// 避免操作时间过长,自动过期时再删除返回结果为0
		if msg == DEL_LOCK_SUCCESS || msg == DEL_LOCK_NON_EXISTENT {
			return true
		}
		return false
	}
	return false
}

// func main() {

//     uid, _ := uuid.NewUUID()
//     id := uid.String()

//     key := "用户id001_操作del_操作方法名DelComment"

//     if pool.AddLock(key, id, 3) {
//         fmt.Println("添加lock成功")
//     } else {
//         fmt.Println("添加lock失败")
//     }

//     if pool.GetLock(key) == id {
//         fmt.Println("相等")
//     }else {
//         fmt.Println("不相等")
//     }

//     if pool.DelLock(key,id){
//         fmt.Println("删除成功")
//     } else {
//         fmt.Println("删除失败")
//     }
// }
