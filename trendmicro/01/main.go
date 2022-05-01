package main

import (
	"fmt"
	"strings"
)

// 回复被攻击的url
// 攻击的方法：
// 1. // 变为\//
// 2. // 变为\\\\
// 把协议://服务器xxxxx   协议和服务器中间的斜杠恢复成正常的

func main() {
	var url string  // https://\\www.baidu.com/////
	url = "https://\\www.baidu.com/////"
	// fmt.Scanln(&url)

	// 分段
	arr := strings.Split(url,":")
	res := arr[0] +":"
	// 处理arr[1]  
	var i int
	bs := []byte(arr[1])
	for i=0;;i++{
		// if bs[i] == '\/' || bs[i] == '\\' {// 问题是转义字符的处理
		if bs[i] == []byte("/")[0] || bs[i] == []byte("\\")[0] {
			continue
		}else{
			break
		}
	}
	res += "//" + string(bs[i:])
	fmt.Println(res)
}