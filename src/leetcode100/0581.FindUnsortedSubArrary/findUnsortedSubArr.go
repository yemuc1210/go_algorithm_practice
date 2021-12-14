package lt581
/*
如果最右端的一部分已经排好序，这部分的每个数都比它左边的最大值要大，
同理，如果最左端的一部分排好序，这每个数都比它右边的最小值小。
所以我们从左往右遍历，如果i位置上的数比它左边部分最大值小，则这个数肯定要排序， 
就这样找到右端不用排序的部分，同理找到左端不用排序的部分，它们之间就是需要排序的部分
*/
func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	length := len(nums)

	max,min := nums[0],nums[len(nums)-1]
	high,low := 0, length-1
	//从左到右，找到最大值，
	for i:=1;i<len(nums);i++{
		max = max1(max,nums[i])
		min = min1(min,nums[length - i - 1])
		if nums[i] < max {
			high = i
		}
		if nums[length-i-1] > min{
			low = length-i-1
		}
	}
	if high > low{
		return high-low+1
	}else{
		return 0
	}
}
func max1(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func min1(a,b int)int{
	if a<b{
		return a
	}
	return b
}