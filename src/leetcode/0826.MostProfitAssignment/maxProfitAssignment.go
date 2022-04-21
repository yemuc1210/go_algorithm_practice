package main

import (
	"fmt"
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	// 难道有坑？贪心不就行了吗
	// 1. 排序
	type pair struct {
		difficulty int
		profit     int
	}
	pairs := make([]pair, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		pairs[i] = pair{difficulty[i], profit[i]}
	}
	// 排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].difficulty < pairs[j].difficulty
	}) // 根据难度升序
	sort.Slice(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})
	// 贪心
	var res int
	for i := 0; i < len(worker); i++ {
		maxProfit := pairs[0].profit
		j := 0
		for j = 0; j < len(pairs) && pairs[j].difficulty <= worker[i]; j++ {
			if pairs[j].profit > maxProfit {
				maxProfit = pairs[j].profit
			}
		}
		if j == 0 {
			maxProfit = 0
		}
		res += maxProfit
	}
	return res
}

func main() {

	type para struct {
		diff   []int
		prof   []int
		worker []int
	}
	paras := []para{
		{
			diff:   []int{2, 4, 6, 8, 19},
			prof:   []int{10, 20, 30, 40, 50},
			worker: []int{4, 5, 6, 7},
		},
		{
			diff:   []int{13, 37, 58},
			prof:   []int{4, 90, 96},
			worker: []int{34, 73, 45},
		},
		{
			[]int{64,88,97},
			[]int{53,86,89},
			[]int{98,11,6},
		},
	}
	for _, pa := range paras {
		res := maxProfitAssignment(pa.diff, pa.prof, pa.worker)
		fmt.Println(res)

	}
}
