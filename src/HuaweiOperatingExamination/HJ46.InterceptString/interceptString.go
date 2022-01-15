package main

import "fmt"


func main() {
	var n int
	var err error
	var s string
	var k int
	for {
		n,err = fmt.Scanln(&s)
		if n==0 || err!=nil{
			break
		}
		fmt.Scanln(&k)
		//截取s前k个字符并输出
		bytes := []byte(s)
		fmt.Println(string(bytes[:k]))
	}
}