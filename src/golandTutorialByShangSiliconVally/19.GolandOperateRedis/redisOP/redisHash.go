package redisop

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)


func RedisHash(){
	//通过golad对redis进行操作

	//1 连接redis
	conn,err := redis.Dial("tcp","127.0.0.1:6379")  //本机 6379端口
	if err != nil {
		//报错
		fmt.Println("redis dial err:",err)
		return
	}

	defer conn.Close()   //关闭

	//2 首先写入 ]
	_,err = conn.Do("hset","user3","name","usr3go","age","23")
	if err != nil {
		fmt.Println("hset err:",err)
		return
	}

	//3 读 hgetall
	res,err := redis.StringMap(conn.Do("hgetall","user3"))
	if err!=nil {
		fmt.Println("hget err:",err)
		return
	}

	//res 是stringMap     ok,hget res: map[age:23 name:usr3go]
	fmt.Println("ok,hget res:",res)


	//另一种读取方法
	//批量读取  hmget
	res1,err1 := redis.Strings(conn.Do("hmget","user3","name","age"))
	if err1!=nil {
		fmt.Println("hmget err:",err)
		return
	}
	for i, v := range res1 {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}