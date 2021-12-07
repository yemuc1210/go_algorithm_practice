package offerS6

/*
剑指 Offer II 006. 排序数组中两个数字之和
给定一个已按照 升序排列  的整数数组 numbers ，
请你从数组中找出两个数满足相加之和等于目标数 target 。

函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。
numbers 的下标 从 0 开始计数 ，所以答案数组应当满足
	 0 <= answer[0] < answer[1] < numbers.length 。

假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。

双指针
*/
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0{
		return []int{}
	}
	if len(numbers)==1 {
		if numbers[0] != target{
			return []int{}
		}else{
			return []int{numbers[0]}
		}  //当然，题目假设了有且只有一对符合条件，所以不会存在单个数字返回的情况
	}
	
	//升序序列  
	frontIdx,rearIdx := 0,len(numbers)-1  //首尾指针

	for frontIdx < rearIdx{
		sum := numbers[frontIdx] + numbers[rearIdx]
		if sum == target{
			return []int{frontIdx,  rearIdx}
		}else if sum>target{
			rearIdx --
		}else{
			frontIdx++
		}
	}
	return []int{}
}