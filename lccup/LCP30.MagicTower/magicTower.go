package lcp30
/*
小扣当前位于魔塔游戏第一层，共有 N 个房间，编号为 0 ~ N-1。
每个房间的补血道具/怪物对于血量影响记于数组 nums，
其中正数表示道具补血数值，即血量增加对应数值；
负数表示怪物造成伤害值，即血量减少对应数值；
0 表示房间对血量无影响。

小扣初始血量为 1，且无上限。
假定小扣原计划按房间编号升序访问所有房间补血/打怪，为保证血量始终为正值，
小扣需对房间访问顺序进行调整，
每次仅能将一个怪物房间（负数的房间）调整至访问顺序末尾。
请返回小扣最少需要调整几次，才能顺利访问所有房间。
若调整顺序也无法访问完全部房间，请返回 -1。


优先权队列

遇见负数先加上再说，同时保存在最小堆里面
如果此时血量不是正数了，那么就从堆中拿出最小的一个放到最后，同时更新血量

*/

// import (
//     "container/heap"
//     "fmt"
// )
 
type IntHeap []int
 
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func magicTower(nums []int) int {
	queue := IntHeap{}
	n := len(nums)
	cur,sum := 1,1
	res := 0
	for i:=0;i<n;i++{
		sum += nums[i]

		if nums[i]<0 {
			queue.Push(nums[i])
			cur += nums[i]
			if cur <= 0{
				val := queue.Pop()
				cur -= val.(int)
				res ++
			}else{
				cur += nums[i]
			}
		}
	}
	if sum >0{
		return res
	}
	return -1
}