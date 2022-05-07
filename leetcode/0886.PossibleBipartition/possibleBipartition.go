package main

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int,n+1)
	for i:=range graph{
		graph[i] = []int{}
	}
	// 着色问题
	for _,dislike := range dislikes{
		graph[dislike[0]] = append(graph[dislike[0]], dislike[1])
		graph[dislike[1]] = append(graph[dislike[1]], dislike[0])
	}
	color := make(map[int]int)
	var dfs  func(node int,c int) bool 
	dfs = func(node int,c int) bool {
		if _,exist:=color[node];exist {
			return color[node] == c
		}
		color[node] = c
		for _,neighbor := range graph[node] {
			if !dfs(neighbor,c^1) {
				return false
			}
		}
		return true
	}
	for node:=1;node<=n;node++{
		if _,ok:=color[node];!ok && !dfs(node,0) {
			return false
		}
	}
	return true
}