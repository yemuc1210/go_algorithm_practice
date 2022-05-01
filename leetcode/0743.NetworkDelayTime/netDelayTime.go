package lt743

import "math"

/*
时间列表time  time[i]=(ui,vi,wi)   权重wi为时间
从某个节点发出信号，需要多久，所有节点都收到

dfs
弗洛伊德算法：各个顶点间的最短路径
*/
// func networkDelayTime(times [][]int, n int, k int) int {
// 	graph := make([][]int,n+1)
// 	for i:=0;i<len(graph);i++{
// 		graph[i] = make([]int, n+1)
// 		//初始化为极大值
// 		for j:=0;j<len(graph[i]);j++{
// 			graph[i][j] = math.MaxInt64
// 		}
// 	}
// 	for i:=1;i<=n;i++{
// 		graph[i][i] = 0
// 	}
// 	for _,v := range times{
// 		graph[v[0]][v[1]] = v[2]
// 	}

// 	for m:=1;k<=n;k++{
// 		for i:=1;i<=n;i++{
// 			for j:=1;j<=n;j++{
// 				graph[i][j] = min(graph[i][j],graph[i][m]+graph[m][j])
// 			}
// 		}
// 	}
// 	res := 0
// 	for i:=1;i<=n;i++{
// 		if graph[k][i] == math.MaxInt64 {
// 			return -1
// 		}
// 		res = max(res,graph[k][i])
// 	}
// 	return res
// }

// func min(a,b int)int{
// 	if a<b{
// 		return a
// 	}
// 	return b
// }
// func max(a,b int)int{
// 	if a>b{
// 		return a
// 	}
// 	return b
// }
func networkDelayTime(times [][]int, n, k int) (ans int) {
    const inf = math.MaxInt64 / 2
    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, n)
        for j := range g[i] {
            g[i][j] = inf
        }
    }
    for _, t := range times {
        x, y := t[0]-1, t[1]-1
        g[x][y] = t[2]
    }

    dist := make([]int, n)
    for i := range dist {
        dist[i] = inf
    }
    dist[k-1] = 0
    used := make([]bool, n)
    for i := 0; i < n; i++ {
        x := -1
        for y, u := range used {
            if !u && (x == -1 || dist[y] < dist[x]) {
                x = y
            }
        }
        used[x] = true
        for y, time := range g[x] {
            dist[y] = min(dist[y], dist[x]+time)
        }
    }

    for _, d := range dist {
        if d == inf {
            return -1
        }
        ans = max(ans, d)
    }
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
