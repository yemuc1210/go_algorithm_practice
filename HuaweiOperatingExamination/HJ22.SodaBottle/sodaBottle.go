package main

import "fmt"


func main(){
	var N int
	
	for {
		n,_ := fmt.Scan(&N)
		if n==0 {  //输入0表示输入结束
			break
		}else{
			if N==0 {
				break
			}
			//空瓶书N
			fmt.Println(N/2)   //其实是两个换一个 每次都借
		}
	}
}
