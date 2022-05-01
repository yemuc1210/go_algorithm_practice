package lt179

import (
	"sort"
	"strconv"
)

/*
179. 最大数
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"
*/
/*
规则就是：第一个数最大的放前面  s1+s2和s2+s11
定义新的排序方法，即less方法
*/
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		//高位比较
		x,y := nums[i],nums[j]
		sx,sy := 10,10
		for sx <= x{
			sx *= 10
		}
		for sy <= y{
			sy *= 10  //得到最高位阶数
		}
		return sy*x+y >sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	res := []byte{}
	for _,num := range nums{
		res = append(res, strconv.Itoa(num)...)
	}
	return string(res)
}