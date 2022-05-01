package main

import "fmt"

//刷墙
/*
最近小明搬到了新家，他正在粉刷墙壁，但是不幸的是他粉刷的墙壁并不理想。
他的墙壁是一个长度为 的格子，每个格子用0表示红色，用1表示蓝色。
现在墙壁是一个非常混乱的颜色。他想将墙壁涂成左边全是蓝色右边全是红色，
可以将墙壁刷成全是红色或者蓝色。请问他至少需要粉刷多少个格子墙壁刷成他想要的样子？
*/
func main() {
	var n int 
	fmt.Scanln(&n)  
	var s string
	fmt.Scanln(&s)  //第二行 01串   0-红 1-蓝
	colors := []byte(s)
	//最少粉刷次数  dp?
	// 左边全是蓝  右边全是红   检索左边红色数目c1+右边蓝色数目c2
	// 全部是红色或蓝色   
	// c1<n/2 && c2<n/2 不超过半数 方案1  
	if n == 0 {
		fmt.Println(0)
		return 
	}
	if n==1 {
		fmt.Println(0)
		return 
	}
	
	c1 := searchCnt(colors[:n/2], '0')
	c2 := searchCnt(colors[n/2:], '1')
	//方案1
	w1 := c1+c2
	// 方案2
	w2 := way2(c1,c2,n)
	if w1 < w2 {
		fmt.Println(w1)
	}else{
		fmt.Println(w2)
	}
}
func way2(c1,c2,n int) int {
	// 全部是红色 或 全部是 蓝色
	// 全部是红色
	r1 := n/2 - c1  + c2
	// 全部是蓝色
	r2 := c1 + n/2 - c2
	if r1 < r2 {
		return r1
	}else{
		return r2
	}
}
func searchCnt(colors []byte, target byte) int {
	res := 0
	// 检索target个数 
	for i:=0;i<len(colors);i++ {
		if colors[i] == target{
			res++
		}
	}
	return res
}