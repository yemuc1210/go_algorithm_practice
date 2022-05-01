package lt611

import "sort"

func triangleNumber(nums []int) int {
	if len(nums) < 3{
		return 0
	}
	//先排序
	sort.Ints(nums)   //升序排
	//此时三边a<=b<=c   只需判断a+b>c即可
	ans :=0
	for i,v := range nums {
		for j:=i+1;j<len(nums);j++{
			ans  += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return ans
}