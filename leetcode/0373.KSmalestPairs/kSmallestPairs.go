package lt373


/*
剑指 Offer II 061. 和最小的 k 个数对
给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
*/
import (
    "container/heap"
)
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
    m, n := len(nums1), len(nums2)
    h := hp{nil, nums1, nums2}
    for i := 0; i < k && i < m; i++ {
        h.data = append(h.data, pair{i, 0})
    }
    for h.Len() > 0 && len(ans) < k {
        p := heap.Pop(&h).(pair)
        i, j := p.i, p.j
        ans = append(ans, []int{nums1[i], nums2[j]})
        if j+1 < n {
            heap.Push(&h, pair{i, j + 1})
        }
    }
    return
}

type pair struct{ i, j int }
type hp struct {
    data         []pair
    nums1, nums2 []int
}
func (h hp) Len() int            { return len(h.data) }
func (h hp) Less(i, j int) bool  { a, b := h.data[i], h.data[j]; return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j] }
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }
