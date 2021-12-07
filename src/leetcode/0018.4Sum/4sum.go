package lt18

import "sort"

/*
18. 四数之和
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
（若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案
*/

/*
不重复  先排序可解决
*/
func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0{
		return [][]int{}
	}
	sort.Ints(nums)
	return nSum(nums,4,0,target)
}


func nSum(nums []int, n int , start int ,target int) [][]int{
	size := len(nums)
	res := [][]int{}

	if n<2 || size<n {
		return res
	}

	if n==2 {  //两数和 用双指针
		left,right := start, size-1
		for left < right {
			leftElem,rightElem := nums[left],nums[right]
			sum := nums[left] + nums[right] 
			if sum < target {
				//左指针右移动 需要去重
				for left<right && nums[left]==leftElem {
					left ++
				}
			}else if sum > target {
				//右指针左移
				for left<right && nums[right]==rightElem {
					right--
				}
			}else{
				//得到结果
				res = append(res, []int{nums[left],nums[right]})
				//去重
				//左指针右移动 需要去重
				for left<right && nums[left]==leftElem {
					left ++
				}
				//右指针左移
				for left<right && nums[right]==rightElem {
					right--
				}
			}
		}
	} else {  //3数和 4数和。。。
		//递归计算
		for i:=start;i<size;i++{
			subRes := nSum(nums,n-1, i+1, target-nums[i])
			//加入当前元素
			for _,list := range subRes {
				res = append(res, append([]int{nums[i]}, list...))
			}
			//去重
			for i<size-1 && nums[i+1] == nums[i] {
				i++
			}
		}
	}
	return res
}