package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

//实现一个简单的tcp 服务端
//tcp粘包测试
func main(){
	//首先，设置监听的端口 即服务端开放的端口
	listener,err  := net.Listen("tcp","127.0.0.1:30000")  
	if err != nil {
		fmt.Println("error ")
		return 
	}
	defer listener.Close()
	for {
		conn,errAcc:= listener.Accept()
		if errAcc != nil {
			fmt.Println("建立连接出错")
			return 
		}
		//根据连接 开启一个协程处理
		go process(conn)
	}
	
}

func process(conn net.Conn) {
	defer conn.Close()
	//读数据
	reader := bufio.NewReader(conn)
	var buf [128]byte
	for {
		len,err := reader.Read(buf[:])
		if err == io.EOF {
			// fmt.Println("服务端读取出错")
			break
		}
		if err!=nil{
			fmt.Println("读取出错")
			break
		}
		//输出数据
		revStr := string(buf[:len])
		fmt.Printf("收到的数据：%s \n",revStr)
	}
	
}