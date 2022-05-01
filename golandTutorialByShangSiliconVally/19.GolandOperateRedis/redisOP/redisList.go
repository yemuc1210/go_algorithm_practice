package redisop

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/*
核心代码

_, err = c.Do("lpush", "heroList", "no1:宋江", 30, "no2:卢俊义", 28)
r, err := redis.String(c.Do("rpop", "heroList"))

*/
func RedisList(){
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..
	//2. 写入
	_,err = conn.Do("lpush","heroList","no1:宋江", 30, "no2:卢俊义", 28)
	if err != nil {
		fmt.Println("lpush err=", err)
		return
	}

	//3 读取
	r, err := redis.String(conn.Do("rpop", "heroList"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	fmt.Println("ok 读取结果是：",r)
}