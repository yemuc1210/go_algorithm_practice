package main

import (
	"fmt"
	"net"
)

//实现一个 tcp客户端
func main(){

	//首先 拨号建立连接
	conn,err := net.Dial("tcp","127.0.0.1:30000")
	if err!=nil {
		fmt.Println("dial fail")
		return 
	}
	defer conn.Close()
	for i:=0;i<10;i++{
		//发送
		msg := fmt.Sprintf("hello %d\n",i)
		conn.Write([]byte(msg))
	}
}