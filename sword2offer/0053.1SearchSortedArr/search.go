package offer53_1
//统计一个数字在排序数组中出现的次数
func search(nums []int, target int) int {
	cnt := 0 
	for _,v := range nums{
		if v == target{
			cnt++
		}
	}
	return cnt
}