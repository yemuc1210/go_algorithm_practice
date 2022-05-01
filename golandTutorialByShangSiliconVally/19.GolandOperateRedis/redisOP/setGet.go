package redisop

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/*
set/get 接口
*/
func RedisSetGet(){
	//通过go 向redis 写入数据和读取数据
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..
	//2. 通过go 向redis 写入数据string [key-val]
	_, err = conn.Do("Set", "name", "tomjerry 猫猫")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//设置有效时间
	// _, err = conn.Do("expire", "name", 10)

	/*批量写入
	_, err = c.Do("MSet", "name", "尚硅谷", "address", "北京昌平~")
	r, err := redis.Strings(c.Do("MGet", "name", "address"))
	for _, v := range r {
		fmt.Println(v)
	}
	*/

	//3. 通过go 向redis 读取数据string [key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	//因为返回r 是interface{}
	//因为name 对应的值是string ,因此我们需要转换
	//nameString := r.(string)
	fmt.Println("操作ok ", r)
}