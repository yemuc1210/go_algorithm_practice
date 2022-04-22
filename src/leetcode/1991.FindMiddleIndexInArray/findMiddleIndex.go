package main

// 找到数组的中间位置
// 中间位置：左右两边和相等
func findMiddleIndex(nums []int) int {
	// 前缀和
	total := 0
	for _,v := range nums {
		total += v
	}
	// 遍历到i个元素时，左侧元素之和为sUm，右侧元素之和为total-nums[i]-sum
	// 左右侧元素相等，sum = total-nums[i]-sum  --->  2sum+nums[i] = total
	// 当中心索引左侧或右侧没有元素时，即为零项相加，【空和】=0
	sum := 0
	for i:=0;i<len(nums);i++{
		if 2*sum+nums[i] == total {
			return i
		}
		sum += nums[i]
	}
	return -1
}

func main(){

}