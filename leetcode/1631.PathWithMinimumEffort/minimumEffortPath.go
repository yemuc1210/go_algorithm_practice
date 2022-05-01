package main

import (
	"container/heap"
	"math"
)

// 路径耗费的体力值 = 路径上 相邻格子高度查绝对值的最大值 决定
// 找到耗费体力最小的路径
// 贪心：每一步都根据贪心去选择，则最终结果
// 记录路径的最大值 最小值，从而得到最大高差

// 并查集？

// dfs 搜索路径 得到所有得路径和对应得最大高差
// 然后从n多个路径中，挑选出最大高差最小的， 可以使用优先权队列
func minimumEffortPath(heights [][]int) int {
	// 定义存储结构，将路径-最大高差联系起来
	h := &IPriorityQueue{}
	heap.Init(h)
	m := make(map[int][]int)
	rows,cols := len(heights),len(heights[0])
	dfs(heights,rows,cols,0,0,[]int{},m,h,math.MaxInt32,math.MinInt32)
	
	return h.Pop().(int)
}

func dfs(heights [][]int, rows,cols int, i,j int, path []int, res map[int][]int,h *IPriorityQueue,curMinHeight int, curMaxHeight int) {
	if i<0 || i>rows || j<0 ||j>cols { // 边界条件
		return  
	}

	if i==rows && j==cols {
		// 得到一组结果
		tmp := make([]int, len(path))
		copy(tmp,path)
		h.Push(curMaxHeight - curMinHeight)
		res[curMaxHeight-curMinHeight] = tmp 
		return 
	}
	// 访问i,j
	if heights[i][j] < curMinHeight {
		curMinHeight = heights[i][j]
	}else if heights[i][j] > curMaxHeight {
		curMaxHeight = heights[i][j]
	}
	path = append(path, heights[i][j])     
	// 四个方向
	dfs(heights,rows,cols,i-1,j,path,res,h,curMinHeight,curMaxHeight)
	dfs(heights,rows,cols,i+1,j,path,res,h,curMinHeight,curMaxHeight)
	dfs(heights,rows,cols,i,j-1,path,res,h,curMinHeight,curMaxHeight)
	dfs(heights,rows,cols,i,j+1,path,res,h,curMinHeight,curMaxHeight)
	path = path[:len(path)-1]
}

////////////////////////////////////////////////////
// 定义优先权队列   小根堆
type IPriorityQueue []int

// 实现接口
func (q IPriorityQueue) Less(i,j int ) bool {
	return q[i] < q[j]   // 小根堆
}
func (q IPriorityQueue) Swap(i,j int) {
	q[i],q[j] = q[j],q[i]
}
func (q IPriorityQueue) Len() int {return len(q)}

// 实现Push Pop
func (h *IPriorityQueue) Push(x interface{}) {
	// 转int
	*h = append(*h, x.(int))
}
func (h *IPriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]  // 移除最后一个
	return x
}
////////////////////////////////////////////////
func main() {

}