package lt528

import (
	"math/rand"
	"sort"
)

/*
设数组 ww 的权重之和为 \textit{total}total。根据题目的要求，我们可以看成将 [1,total] 范围内的所有整数分成 n 个部分（其中 n 是数组 w 的长度），
第 i 个部分恰好包含 w[i] 个整数，并且这 n 个部分两两的交集为空。
随后我们在 [1,total] 范围内随机选择一个整数 x，如果整数 x 被包含在第 i 个部分内，我们就返回 i。

一种较为简单的划分方法是按照从小到大的顺序依次划分每个部分。
例如 w=[3,1,2,4] 时，权重之和 total=10，那么我们按照[1,3],[4,4],[5,6],[7,10] 对 [1, 10][1,10] 进行划分，
使得它们的长度恰好依次为 3,1,2,4。
可以发现，每个区间的左边界是在它之前出现的所有元素的和加上 1，右边界是到它为止的所有元素的和。
因此，如果我们用 pre[i] 表示数组 ww 的前缀和：

\textit{pre}[i] = \sum_{k=0}^i w[k]

那么第 i 个区间的左边界就是 pre[i]−w[i]+1，右边界就是 pre[i]。

当划分完成后，假设我们随机到了整数 x，我们希望找到满足：

pre[i]−w[i]+1≤x≤pre[i]
的 i 并将其作为答案返回。由于 pre[i] 是单调递增的，
因此我们可以使用二分查找在O(logn) 的时间内快速找到 i，即找出最小的满足 x≤pre[i] 的下标 i。
*/
type Solution struct {
	prefix []int
}


func Constructor(w []int) Solution {
	for i:=1;i<len(w);i++{
		w[i] += w[i-1]
	}
	return Solution{prefix: w}
}


func (this *Solution) PickIndex() int {
	x := rand.Intn(this.prefix[len(this.prefix)-1]) + 1
	return sort.SearchInts(this.prefix, x)
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */