package lt268
/*
268. 丢失的数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，
找出 [0, n] 这个范围内没有出现在数组中的那个数。 
*/
func missingNumber(nums []int) int {
	//只丢失一个数  直接求和
	sum := 0
	for _,v := range nums{
		sum += v
	}
	return (1+len(nums))*len(nums)/2 - sum
}