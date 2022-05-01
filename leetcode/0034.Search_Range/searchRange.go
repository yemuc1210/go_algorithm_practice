package lt34
/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。
找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：
	你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

那就是二分搜索
先实现
*/
func searchRange(nums []int, target int) []int {
    if len(nums)==0{
        return []int{-1,-1}
    }
	start,end := 0,0
	if nums[0]==target{
		end = start+1
	
		for end<len(nums)&&nums[end]==target {end++}
	
		return []int{start,end-1}
	}
	for i:=end;i<len(nums);{
		if nums[i]==target{
			start = i		
			j:=i
			for ;j<len(nums)&&nums[j]==target;j++{}
			end = j-1
			i = j
			return []int{start,end}
		}else{
			i++
		}
	}

	return []int{-1,-1}
}