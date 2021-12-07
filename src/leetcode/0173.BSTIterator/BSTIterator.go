package lt173
import (
	"structs"
	"container/heap"
)
type TreeNode = structs.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 使用优先队列
 */
type BSTIterator struct {
	pq PriorityQueueOfInt
	count int
}


func Constructor(root *TreeNode) BSTIterator {
	result, pq := []int{},PriorityQueueOfInt{}
	postorder(root, &result)
	for _, v := range result {
		heap.Push(&pq, v)
	}
	bs := BSTIterator{pq: pq, count: len(result)}
	return bs
}

func postorder(root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		*output = append(*output, root.Val)
	}
}

/** @return the next smallest number */
func (bs *BSTIterator) Next() int {
	bs.count--
	return heap.Pop(&bs.pq).(int)
}

/** @return whether we have a next smallest number */
func (bs *BSTIterator) HasNext() bool {
	return bs.count != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
 type PriorityQueueOfInt []int

func (pq PriorityQueueOfInt) Len() int {
	return len(pq)
}

func (pq PriorityQueueOfInt) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueueOfInt) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueOfInt) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *PriorityQueueOfInt) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
