package main

import "fmt"

//找零
//1 4 16 64（硬币）     1024
//最少收到多少硬币
func main() {
	var N int
	fmt.Scanln(&N)
	N = 1024-N   
	//求N最少等于多少枚硬币
	//贪心法
	coins := 64
	cnt := 0
	for i:=0;i<4;i++ {
		cnt += N/coins
		N %= coins
		coins >>= 2  // 换下一个面值
	}	
	fmt.Println(cnt)
}