package NC82

import (
	"container/heap"
	"sort"
)

/**
 * 给定窗口尺寸，找出所有窗口中最大值
 * 剑指offer59  lt239
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
 */
var a []int
type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func maxSlidingWindow(nums []int, k int) []int {
    a = nums
    q := &hp{make([]int, k)}
    for i := 0; i < k; i++ {
        q.IntSlice[i] = i
    }
    heap.Init(q)

    n := len(nums)
    ans := make([]int, 1, n-k+1)
    ans[0] = nums[q.IntSlice[0]]
    for i := k; i < n; i++ {
        heap.Push(q, i)
        for q.IntSlice[0] <= i-k {
            heap.Pop(q)
        }
        ans = append(ans, nums[q.IntSlice[0]])
    }
    return ans
}
func maxInWindow(nums []int, k int) []int {
    q := []int{}
    push := func(i int) {
        for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
            q = q[:len(q)-1]
        }
        q = append(q, i)
    }

    for i := 0; i < k; i++ {
        push(i)
    }

    n := len(nums)
    ans := make([]int, 1, n-k+1)
    ans[0] = nums[q[0]]
    for i := k; i < n; i++ {
        push(i)
        for q[0] <= i-k {
            q = q[1:]
        }
        ans = append(ans, nums[q[0]])
    }
    return ans
}

/**
 * 可能长度为0 k>length等情况
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
*/
func maxInWindows1(num []int, size int) []int {
    // write code here
    l := len(num)
 
    if l == 0 || size == 0 || size > l {
        return nil
    }
 
//  返回的结果集
    res := make([]int, 0)
//  单调队列 
//  从大到小排序
    list := make([]int, 0)
 
//  窗口右移，判断单调队列第一个元素是否等于被移除的元素值
//  等于被移除的元素值，则单调队列第一个元素也移除
    pop := func(val int) {
        if list[0] == val {
            list = list[1:]
        }
    }
 
//  向单调队列压入一个数据
//  单调队列:  大于等于左边元素  小于等于右边元素
    push := func(val int) {
        for len(list) != 0 && list[len(list)-1] < val {
            list = list[:len(list)-1]
        }
 
        list = append(list, val)
    }
 
//  前size个元素进入单调队列
    for i := 0; i < size; i++ {
        push(num[i])
    }
//  前size个元素为第一个窗口，取第一个窗口最大值
    res = append(res, list[0])
 
//  处理后续元素
    for i := size; i < len(num); i++ {
//      处理窗口左侧被移除的元素，判断是否也需要从单调队列移除
        pop(num[i-size])
//      添加当前窗口右侧新增的元素
        push(num[i])
 
//      取单调队列最大的元素值
        res = append(res, list[0])
    }
 
    return res
}