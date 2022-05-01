package lt347

import "container/heap"
/*

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
把数组构造成一个优先队列，输出前 K 个即可。
*/

type Item struct{
	key,count int
}

//实现优先权队列
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int{
	return len(pq)
}

func (pq PriorityQueue) Less(i,j int) bool{
	// golang中的heap是按最小堆组织的，所以count越大，Less()越小，越靠近堆顶.
	return pq[i].count > pq[j].count
}
func (pq PriorityQueue) Swap(i,j int){
	pq[i],pq[j] = pq[j],pq[i]
}

func (pq *PriorityQueue) Push(x interface{}){
	item := x.(*Item)
	*pq = append(*pq, item) 
}

func (pq *PriorityQueue) Pop() interface{}{
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m:=make(map[int]int)
	for _,n := range nums{
		m[n]++
	}
	pq := PriorityQueue{}
	for key,count := range m{
		heap.Push(&pq, &Item{key: key, count: count})
	}
	var result []int
	for len(result) < k{
		item := heap.Pop(&pq).(*Item)
		result = append(result, item.key)
	}
	return result

}