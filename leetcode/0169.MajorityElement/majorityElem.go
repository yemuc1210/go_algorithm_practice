package lt169

import "sort"

/*
给定一个大小为 n 的数组，找到其中的多数元素。
多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

*/
func majorityElement(nums []int) int {
	//排序  然后返回中心值
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//进阶，尝试时间复杂度O(n),空间复杂度O(1)
//moors 投票，遇到相同的+1 不同的-1  为0换数
func majorityElement1(nums []int)int{
	cur,cnt := nums[0],0
	for i:=0;i<len(nums);i++{
		if cnt == 0{
			cur,cnt = nums[i],0 
		}else{
			if nums[i] == cur{
				cnt+=1
			}else{
				cnt -=1
			}
		}
	}
	return cur
}