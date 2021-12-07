package interview17

import (
	"math/rand"
	"time"
)

/*
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

基于快排，快排每次会划分两部分，大于的在右小于的在左
这里只划分一边
使得前k个在数组左侧


*/
func smallestK(arr []int, k int) []int {
	//使用快排
	if k == 0 {
        return nil
    }
    rand.Seed(time.Now().UnixNano())
    randomizedSelected(arr, 0, len(arr)-1, k)
    return arr[:k]

}

func partition(nums []int, l,r int) int{
	pivot := nums[r]

	i := l-1
	for j:=l; j<r; j++ {
		if nums[j] <= pivot {
			i++
			nums[j],nums[i] = nums[i],nums[j]
		}
	}
	nums[i+1],nums[r] = nums[r],nums[i+1]
	return i+1
}

func randomPartition(nums []int, l,r int) int{
	i := l+rand.Intn(r-l+1)
	nums[i],nums[r] = nums[r],nums[i]
	return partition(nums,l,r)
}
func randomizedSelected(arr []int, l, r, k int) {
    pos := randomPartition(arr, l, r)
    num := pos - l + 1
    if k < num {
        randomizedSelected(arr, l, pos-1, k)
    } else if k > num {
        randomizedSelected(arr, pos+1, r, k-num)
    }
}

