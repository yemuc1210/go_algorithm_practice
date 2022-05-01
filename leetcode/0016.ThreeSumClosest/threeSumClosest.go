package lt16
import (
	"math"
	"sort"
)

/*
16. 最接近的三数之和
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。
请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。
*/
/*
用两个指针夹逼的方法。
先对数组进行排序，i 从头开始往后面扫。
这里同样需要注意数组中存在多个重复数字的问题。
具体处理方法很多，可以用 map 计数去重。
	i 在循环的时候和前一个数进行比较，
	如果相等，i 继续往后移，直到移到下一个和前一个数字不同的位置。
	j，k 两个指针开始一前一后夹逼。j 为 i 的下一个数字，k 为数组最后一个数字，
	由于经过排序，所以 k 的数字最大。
	j 往后移动，k 往前移动，逐渐夹逼出最接近 target 的值。


用暴力解法，三层循环找到距离 target 最近的组合
*/
func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func abs(a int) int{
	if a>=0 {
		return a
	}
	return -a
}
