package main

import "fmt"

type item struct {
	word string
	cnt  int
}

func topKFrequent(words []string, k int) []string {
	// 统计出现次数，使用map
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	// 根据字典序排列
	items := []item{}
	for k, v := range m {
		it := item{k, v}
		items = append(items, it)
	}
	buildHeap(items, len(items))
	fmt.Println(items)
	// 取值
	heapSize := len(items)
	var res []string
	// 4-2   3-3
	for i := len(items) - 1; i >= len(items)-k; i-- {
		items[0], items[i] = items[i], items[0]
		res = append(res, items[i].word)
		heapSize--
		heapMaxify(items, 0, heapSize)
	}
	return res
}

func buildHeap(items []item, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		heapMaxify(items, i, heapSize)
	}
}
func heapMaxify(items []item, i int, heapSize int) {
	// 以i为根节点，需要使其符合最大堆
	l, r := 2*i+1, 2*i+2
	largerOne := i
	if l < heapSize && items[l].cnt > items[largerOne].cnt {
		largerOne = l
	}
	if r < heapSize && items[r].cnt > items[largerOne].cnt {
		largerOne = r
	}
	if largerOne != i {
		// 需要调整
		// 交换
		items[largerOne], items[i] = items[i], items[largerOne]
		heapMaxify(items, largerOne, heapSize)
	}
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	res := topKFrequent(words, k)
	fmt.Println(res)
}
