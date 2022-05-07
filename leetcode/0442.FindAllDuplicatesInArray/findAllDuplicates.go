package main

func findDuplicates(nums []int) []int {
	res := []int{}
	// 数字范围1-n  下标范围0-n-1
	for _,v := range nums {
		if v<0 {
			v = -v
		}
		// nums[i] －  表示数字i+1出现过
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}else{
			res = append(res, v)
		}
	}
	return res
}