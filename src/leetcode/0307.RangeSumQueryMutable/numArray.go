package main

import "math"

/*
数组需要完成
- 要求更新数组下标对应的值
- 返回数组中索引的元素和  双闭区间
*/

// 分块
type NumArray struct {
	nums, sums []int // sums[i] 表示第 i 个块的元素和
	size       int   // 块的大小
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	size := int(math.Sqrt(float64(n)))
	sums := make([]int, (n+size-1)/size) // n/size 向上取整
	for i, num := range nums {
		sums[i/size] += num
	}
	return NumArray{nums, sums, size}
}

func (na *NumArray) Update(index, val int) {
	na.sums[index/na.size] += val - na.nums[index]
	na.nums[index] = val
}

func (na *NumArray) SumRange(left, right int) (ans int) {
	size := na.size
	b1, b2 := left/size, right/size
	if b1 == b2 { // 区间 [left, right] 在同一块中
		for i := left; i <= right; i++ {
			ans += na.nums[i]
		}
		return
	}
	// 不再同一块  则总和包括三个部分  【i，size-1)  [0,i2] 中间块
	for i := left; i < (b1+1)*size; i++ {
		ans += na.nums[i]
	}
	for i := b1 + 1; i < b2; i++ {
		ans += na.sums[i]
	}
	for i := b2 * size; i <= right; i++ {
		ans += na.nums[i]
	}
	return
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
