package lcp28

import "sort"

/*
LCP 28. 采购方案
小力将 N 个零件的报价存于数组 nums。小力预算为 target，
假定小力仅购买两个零件，要求购买零件的花费不超过预算，请问他有多少种采购方案。

注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1


排序+双指针
*/
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)

	mod := 1000000007
	res := 0
	for left,right := 0,len(nums)-1; left < right; left++ {
		for left<right && nums[left]+nums[right] > target {
			right--
		}
		if left<right {
			res = res + right - left
			res = res % mod
		}
	}
	return res

}