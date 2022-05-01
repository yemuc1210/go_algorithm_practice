package offerS12

/*  lt724
剑指 Offer II 012. 左右两边子数组的和相等
给你一个整数数组 nums ，请计算数组的 中心下标 。

数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。

如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。

如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。


记数组的全部元素之和为total，当遍历到第 i 个元素时，设其左侧元素之和为 sum，
则其右侧元素之和为total−numsi−sum。
左右侧元素相等即为 sum=total−numsi-sum
2×sum+numsi=total。

当中心索引左侧或右侧没有元素时，即为零个项相加，这在数学上称作「空和」
。在程序设计中我们约定「空和是零」。

*/
func pivotIndex(nums []int) int {
	if len(nums) == 0 || len(nums) == 1{
		return 0
	}

	total := 0
	for _,v := range nums{
		total += v
	}
	//根据公式找sum
	sum := 0
	for i,v := range nums{
		if 2*sum+v == total{
			return i
		}
		sum += v
	}

	return -1
}