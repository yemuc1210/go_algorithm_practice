package main

import "fmt"

// 你和一群强盗准备打劫银行。给你一个下标从 0 开始的整数数组 security ，
// 其中 security[i] 是第 i 天执勤警卫的数量。
// 日子从 0 开始编号。同时给你一个整数 time 。
// 如果第 i 天满足以下所有条件，我们称它为一个适合打劫银行的日子：
// 第 i 天前和后都分别至少有 time 天。
// 第 i 天前连续 time 天警卫数目都是非递增的。
// 第 i 天后连续 time 天警卫数目都是非递减的。
func goodDaysToRobBank(security []int, time int) []int {
	// 意思是time天的警卫人数是个极小点
	if len(security) < 2*time {
		return nil
	}
	// 找出候选 
	n := len(security)
	// candidates := security[time:n-time]
	// fmt.Println(candidates)
	var res []int	

	// 模拟可以做
	// 单调栈 单调减  单调增 
	// 如果有连续值 如单调减0 1，单调增 3 4，则找到一个2
	// 预处理 表示当前security与前一个大小 1 -1 0
	g := make([]int,n)
	for i:=1;i<n;i++{
		if security[i]>security[i-1] {
			g[i] = 1
		}else if security[i]<security[i-1]{
			g[i] = -1
		}else{
			g[i] = 0
		}
	}
	// 对g使用前缀和的思路
	// a记录前缀1的数量 b记录前缀-1
	// 若i适合打劫，则 i位于candidates && (i-time,i]范围内前缀1为0，(i,i+time]前缀-1为0
	a,b := make([]int,n+1),make([]int,n+1)
	for i:=1;i<=n;i++ {
		// a[i] = a[i-1] + (g[i-1]==1 ?1:0)
		if g[i-1] == 1 {
			a[i] = a[i-1] + 1
		}else{
			a[i] = a[i-1] 
		}
	}
	for i:=1;i<=n;i++ {
		// a[i] = a[i-1] + (g[i-1]==1 ?1:0)
		if g[i-1] == -1 {
			b[i] = b[i-1] + 1
		}else{
			b[i] = b[i-1] 
		}
	}
	for i:=time;i<n-time;i++{
		// 前缀1个数  (i-time,i]
		c1 := a[i+1] - a[i+1-time]
		c2 := b[i+1+time]-b[i+1]
		if c1==0 && c2==0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	security := []int{5,3,3,3,5,6,2}
	time := 2
	fmt.Println(goodDaysToRobBank(security,time))
}