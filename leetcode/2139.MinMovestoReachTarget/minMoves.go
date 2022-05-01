package main

import "fmt"

// 得到目标值的最少行动次数
// 从1开始，可以选择递增、加倍两个操作，加倍至多使用maxDoubles
// 求得到target的最少行动次数
func minMoves(target int, maxDoubles int) int {
	// 逆向，从target先倍÷
	cnt := 0
	// 1 若target不是偶数，不能先直接整除
	for maxDoubles!=0 && target>1 {
		if target%2!=0 {
			cnt++
			target--
			continue
		}
		target /= 2
		cnt ++
		maxDoubles -- 
	}
	// 此时除法不能再用了
	// 判断target值
	if target == 1 {
		return cnt 
	}
	// 否则 是maxDoubles=0，此时只能递减
	return cnt+target-1
}

func main() {
	target := 19
	maxDoubles := 2
	res := minMoves(target,maxDoubles)
	fmt.Println(res)
}