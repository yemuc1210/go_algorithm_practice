package main

import (
	"fmt"
	// "strings"
)

func main() {
	// var ip string
	var a,b,c,d int
	for {
		n,err:=fmt.Scanf("%d.%d.%d.%d",&a,&b,&c,&d)
		if n==0 {
			break
		}
		if err != nil {
			panic(err)
		}
		//判断ip是否是内网
		if a==10 || a==127 || (a==172 && (b>=16 && b<=31)) || (a==192 && b==168) {
			fmt.Println(1)
		}else{
			fmt.Println(0)
		}
  	}
}

/*
思路：ip转换为整数，判断数字是否属于内网地址对应的数字：
10.0.0.0 - 10.255.255.255:  需要判断：10
127.0.0.0 - 127.255.255.255 需要判断： 127
172.16.0.0 - 172.31.255.255 需要判断： 17216 和17231
192.168.0.0 - 192.168.255.255 需要判断： 192168
*/
