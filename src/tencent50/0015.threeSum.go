package tencent50

import "sort"

/*中
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
/*
排序
先得两个数得和  然后问题转换成两个数为0得问题  用hashmap
双指针
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res,start,end, idx := make([][]int,0),0,0,0
	addNum,length := 0,len(nums)

	for idx=1;idx<length-1;idx++{
		start,end = 0,length-1
		if idx>1 && nums[idx] == nums[idx-1] {
			start = idx - 1
		}
		for start<idx && end > idx {
			if start>0 && nums[start] == nums[start-1] {
				start ++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end -- 
				continue
			}
			addNum = nums[start] + nums[end] +nums[idx]
			if addNum == 0 {
				res = append(res, []int{nums[start],nums[idx],nums[end]})
				start++
				end --
			}else if addNum>0{
				end--
			}else{
				start++
			}

		}
	}
	return res
}