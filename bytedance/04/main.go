package main

import "fmt"


func main() {
	var n int
	fmt.Scanln(&n)
	var values,nums = make([]int,n),make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d%c",&values[i])
	}
	for i:=0;i<n;i++ {
		fmt.Scanf("%d%c", &nums[i])
	}

	// 计算
}
func solve(n int, values []int, nums []int) int {
	// 一个宝石变成另一个  连续区间内翻转  使得种类/价值变化，数量不变
	// 最多施展一次  求总价值最大
	// 方法1 暴力解法 双指针
	return 0
}

func solution1(n int,values,nums []int) int {
	// dp[i] 表示[0,i] 内可以取得的宝石最大总价值
	dp := make([]int,n+1)
	dp[0] = values[0] * nums[0]
	if values[0]*nums[1]+values[1]*nums[0] > values[1]*nums[1]+values[0]*nums[0] {
		dp[1] = values[0]*nums[1]+values[1]*nums[0] 
	}else{
		dp[1] = dp[0] + values[1]*nums[1]
	}
	for i:=1;i<n;i++{

	}
	return 0
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
func reverse(values []int, i,j int) {
	// 逆置【i，j】

}