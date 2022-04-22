package main

import (
	"fmt"
	"math"
)

// 给定长度为n的整数数组nums
// 假设 arrk时nums顺时针旋转k个位置后的数组，定义nums的旋转函数
// 定义nums的旋转函数F为：F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]
// 返回F(0) F(1) ... F(n-1)中的最大值
// 第一种：暴力，依次尝试旋转位置
// 第二种：最值问题，暴力 dp等
// 旋转函数
// 暴力法超时
func maxRotateFunction(nums []int) int {
	// 暴力
	F := make([]int, len(nums))
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		F[i] = calculateFunction(nums, i)
		if F[i] > max {
			max = F[i]
		}
	}
	return max
}

func calculateFunction(nums []int, k int) int {
	// 旋转
	// 1. 截断
	n := len(nums)
	a := nums[:n-k] // 左半部分
	b := nums[n-k:] // 右
	nums = append(b, a...)
	// fmt.Println(k, nums)
	// 计算
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res + i*nums[i]
	}
	return res
}
// F(0) = 0*A[0]+1*A[1]+2*A[2]+3*A[3]
// F(1) = 0*A[3]+1*A[0]+2*A[1]+3*A[2]
// F(2) = 0*A[2]+1*A[3]+2*A[0]+3*A[1]
// F(3) = 0*A[1]+1*A[2]+2*A[3]+3*A[0]

// F(1)-F(0) = A[0]+A[1]+A[2]-3*A[3] = sumA - 4*A(3)
// F(2)-F(1) = A[0]+A[1]+A[3]-3*A[2] = sumA - 4*A(2)
// F(3)-F(2) = A[0]+A[2]+A[3]-3*A[1] = sumA - 4*A(1)
// sumA = A[0]+A[1]+A[2]+A[3]

// len(nums) = n+1    F(k+1)-F(k) = sumA - n*A[n-k-1]
// F(k)=F(k−1)+numSum−n×nums[n−k]
func maxRotateFunction1(nums []int) int{
	// 找规律
	// sum
	sum := 0
	f := 0
	for i,v:=range nums{
		sum+=v
		f = f+i*v
	}
	// 求解
	// 根据相邻的，可以递推
	res := f  // F(0)
	n := len(nums)
	for i:=1;i<len(nums);i++{
		tmp := f + sum - n*nums[n-i]
		if tmp > res {
			res = tmp
		}
		f = tmp
	} 
	return res
}

// 前缀和+滑动窗口
func maxRotateFunction2(nums []int) int {
	n := len(nums)
	sum := make([]int, 2*n+10)
	for i:=1;i<=2*n;i++{
		sum[i] = sum[i-1] + nums[(i-1)%n]
	}
	ans := 0 
	for i:=1;i<=n;i++{
		ans += nums[i-1]*(i-1)
	}
	for i,cur:=n+1,ans; i<2*n;i++{
		cur += nums[(i-1)%n] *(n-1)
		cur -= sum[i-1] - sum[i-n]
		if cur>ans {
			ans = cur
		}
	}
	return ans
}
func main() {
	nums := []int{4, 3, 2, 6}
	res := maxRotateFunction(nums)
	fmt.Println(res)
}
