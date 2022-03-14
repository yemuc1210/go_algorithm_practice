package main

import "fmt"


func main() {
	var n int
	fmt.Scanln(&n)
	targetLocation := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanln(&targetLocation[i])
	}

	// 每次从目标点出发时，初始跳1个单位，之后可跳跃长度+1，且可以选择向左、向右
	// 到达新的目标后 跳跃距离重置
	iniLoc := 0
	
}