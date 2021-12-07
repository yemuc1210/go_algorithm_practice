package lt334

import "math"

/*
找到递增的三元组
a 始终记录最小元素，b 为某个子序列里第二大的数。

接下来不断更新 a，同时保持 b 尽可能的小。

如果下一个元素比 b 大，说明找到了三元组。

关键是比较第二小的数
*/
func increasingTriplet(nums []int) bool {
	a := math.MaxInt64
	b := math.MaxInt64

	for _,e := range nums {
		if e <= a {
			a = e
		} else if e <= b {
			b = e
		} else {
			return true
		}
	}
	return false
}