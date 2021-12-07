package redisop

/*
连接池

*/
import (
    "github.com/garyburd/redigo/redis"
    "fmt"
)

var Pool redis.Pool
func init()  {      //init 用于初始化一些参数，先于main执行
    Pool = redis.Pool{
        MaxIdle:     16,//最大的空闲连接数，即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
        MaxActive:   32,//最大的激活连接数，同时最多有N个连接
        IdleTimeout: 120,//空闲连接等待时间，超过此时间后，空闲连接将被关闭
        Dial: func() (redis.Conn, error) {
            return redis.Dial("tcp", "127.0.0.1:6379")
        },
    }
}

func RedisPool() {
    conn :=Pool.Get()
    res,err := conn.Do("HSET","student","name","jack")
    fmt.Println(res,err)
    res1,err := redis.String(conn.Do("HGET","student","name"))
    fmt.Printf("res:%s,error:%v",res1,err)

}
//0 <nil>
//res:jack,error:<nil>
 

