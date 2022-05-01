package lt713
/*
给定一个正整数数组 nums和整数 k 。

请找出该数组内乘积小于 k 的连续的子数组的个数。

滑动数组题
窗口内每加入一个数，如果乘积仍然小于k，则 满足条件的子数组数量 += 窗口内元素数量

例如 [1.2.3.4.5] , k = 7, 符合要求子数组数量为res=0;

[1] res += 子数组.length; res=1;

[1,2] res =res+ 2 = 3

[1,2,3] res = res+3 =6

当遇到比k还大的元素时，直接重新在这个元素的右边继续开始计算

当窗口内元素乘积大于等于k时，在左边踢出元素，收缩窗口，直到乘积不再大于等于k
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	left,right :=0,0
	mul := 1
	res :=0

	for right < len(nums){
		num := nums[right]
		right ++
		//当前元素大于k，从右边重新开始
		if num >= k{
			left = right
			mul = 1
			continue
		}
		mul = mul * num
		//乘积大于等于k,左边缩小
		for mul >= k {
			mul = mul / nums[left]
			left++
		}
		if mul < k{
			res += (right-left)
		}
	}

	return res
}