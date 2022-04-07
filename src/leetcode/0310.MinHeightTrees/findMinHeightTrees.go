package main

import "fmt"

// 将树中任意节点作为根 所有可能的树中，具有最小高度的树，返回根节点标签列表
// 即，从任意节点出发，找到叶子节点，得到一条路径，这些路径长度最高的是树高
// 找到所有树高里面最低的
func findMinHeightTrees(n int, edges [][]int) []int {
	// n个节点  n-1 个无向边
	var res []int

	// 先构建图吧，使用邻接链表
	// 邻接矩阵最简单
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}

	// 根据度预处理，度最大的显然适合做根，因为会有更多的节点在第二层次
	degrees := make([]int, n) //n个节点的度
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a][b]++
		// graph[b][a]++
		degrees[a]++
		degrees[b]++
	}
	// 得到完整的度数，需要/2，不过目的是得到°最大的结点，÷不是必须的
	// 得到
	maxDegree := degrees[0]
	candidates := []int{}
	for i := 1; i < n; i++ {
		if degrees[i] > maxDegree {
			maxDegree = degrees[i]
			// candidates = append(candidates, i)
		}
	}
	// 得到一组候选
	for i := range degrees {
		if degrees[i] == maxDegree {
			candidates = append(candidates, i)

		}
	}
	fmt.Println(candidates)
	// 从候选开始bfs
	var bfs func(startIdx int, curLen int)
	visited := make([]bool, n)
	minLen := n + 1
	bfs = func(startIdx int, curLen int) {
		q := []int{startIdx}
		visited[startIdx] = true
		for len(q) > 0 {
			curWidth := len(q)
			for i := 0; i < curWidth; i++ {
				cur := q[0]
				q = q[1:]
				visited[cur] = true
				// 遍历边，找到下一层
				for j := 0; j < n; j++ {
					if graph[cur][j] != 0 {
						// 下一层节点
						q = append(q, j)
					}
				}
			}
			curLen++
		}
	}
	type pair struct {
		length int
		node   int
	}
	tmp := []pair{}
	for i := 0; i < len(candidates); i++ {
		len := 0
		bfs(candidates[i], len)
		// 找到最小的
		if len < minLen {
			minLen = len
			tmp = append(tmp, pair{length: len, node: i})
		}
	}
	// 得到res
	for i := 0; i < len(tmp); i++ {
		if tmp[i].length == minLen {
			res = append(res, tmp[i].node)
		}
	}
	return res
}

func main() {
	n:=6
	edges := [][]int{
		{3,0},{3,1},{3,2},{3,4},{5,4},
	}
	res := findMinHeightTrees(n,edges)
	fmt.Println(res)
}