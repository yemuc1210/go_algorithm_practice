package offer57



/*
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
如果有多对数字的和等于s，则输出任意一对即可。


示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]

限制：
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^6
*/
func twoSum(nums []int, target int) []int {
	if len(nums) < 2{
		return []int{}
	}
	// res := []int{}
	m := map[int]int{}
	for _,v := range nums{
		other := target - v
		if m[other] != 0 {
			return []int{v,other}

		}
		m[v] = other
	}

	return []int{}
}