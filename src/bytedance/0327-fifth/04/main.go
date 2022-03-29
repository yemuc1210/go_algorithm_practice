package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	graph1 := make(map[int]map[int]int)
	// graph := make([][]int,n+1)
	var m int
	type ss struct {
		cnt  int
		cur  int
		list []int
	}
	graph := make([]ss, n+1)
	for i := 1; i <= n; i++ {
		// 读取一个数字
		fmt.Scanf("%d%c", &m)
		tmp := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d%c", &tmp[j])
		}
		tmpM := make(map[int]int)
		for j := 0; j < m; j++ {
			tmpM[tmp[j]]++
		}
		s := &ss{m, i, tmp}
		graph[i] = *s
		graph1[i] = tmpM
	}
	// 头子选择从一个间谍开始
	sort.Slice(graph, func(i, j int) bool {
		return graph[i].cnt > graph[j].cnt
	})// [{4 2 [1 3 4 5]} {2 1 [2 5]} {2 3 [2 4]} {2 4 [2 3]} {2 5 [1 2]} {0 0 []}]


	var cnt int
	for i := 0; i <= n && len(graph1)!=0; i++ {
		cur := graph[i].cur
		arr := graph[i].list
		// 通知m中的元素
		for _, v := range arr {
			delete(graph1, v)
		}
		delete(graph1, cur)
		cnt++
	}
	fmt.Println(cnt)
}
