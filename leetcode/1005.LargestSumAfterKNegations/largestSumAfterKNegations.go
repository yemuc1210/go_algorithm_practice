package lt1005

import (
	"math"
	"sort"
)

/*
1005. K 次取反后最大化的数组和
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
	选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。
*/
/*
贪心
有负数修改负数，且优先修改负数中最小的（绝对值更大）
当然 可以重复选择同一个 所以第二选择是修改值为0的
其次才是整数
*/
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)   //先排序
	idx,sum := 0,0
	largestNegative := math.MinInt32
	largestNegativeIdx := -1
	for k > 0{
		//可能存在k比数组长度还大的情况
		idx = idx % len(nums)   
		if nums[idx] < 0 {
			//turn up
			//更新
			if largestNegative < nums[idx] {
				largestNegative = nums[idx]
				largestNegativeIdx = idx
			}
			k -- 
			nums[idx] = 0- nums[idx]
			idx ++
			// largestNegative = max(largestNegative,nums[idx]) 
		}else if nums[idx] == 0 {
			//可以跳出 
			k = 0   //剩下的次数重复选择元素0
			break
		}else{
			//遇到正数  对最小的正数操作即可  k是剩余翻转次数
			//如果k是偶数 则实际上不产生改变
			if k%2 == 0 {
				k = 0
				break
			}else{ //奇数   翻转当前最小的正数
				//比较初始最大负数（现在就可能是最小的正数）和当前元素的绝对值
				if nums[idx] < 0-largestNegative {
					nums[idx] = 0 - nums[idx]
				}else{
					//翻转最小的正数
					nums[largestNegativeIdx] = 0 - nums[largestNegativeIdx]
				}
				k = 0
				break
			}
		}
	}
	//计算sum
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	return sum
}
