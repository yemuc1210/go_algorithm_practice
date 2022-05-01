package lt977

/*
双指针
非递减排序数组
返回每个数的平方组成的新数组，要求新数组仍然非递减排序

利用原数组有序。最终数组的最大值由原数组的最大值或者最小值得到，所以从数组最后一开始排列最终数组，双指针首位
*/

func sortedSquares(nums []int) []int{
	res := make([]int, len(nums))
	for i,k,j := 0,len(nums)-1,len(res)-1; i<=j; k--{
		//i,j  首尾指针   k最终结果的指针
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[k] = nums[i]*nums[i]
			i++
		}else{
			res[k] = nums[j]*nums[j]
			j--
		}
	}
	return res
}