package main

import "fmt"

// 两个数组公共的，长度最长的子数组的长度
func findLength(nums1 []int, nums2 []int) int {
	// 最长公共子序列    区别：子数组默认连续
	// dp[i][j] 表示以nums1[i-1] nums2[j-1] 结尾的公共子数组长度
	// 状态转移方程 dp[i][j] = dp[i-1][j-1] + 1   dp[i][j]=0
	// dp[i][j] 标识 A[i:] B[j:]的最长公共前缀
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if res < dp[i][j] {

				res = dp[i][j]
			}
		}
	}

	return res

}
func main() {
	fmt.Println("vim-go")
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	res := findLength(nums1, nums2)
	fmt.Println(res)

}
