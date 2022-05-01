package interview04
/*
面试题 04.01. 节点间通路

节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。

dfs bfs

*/
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	edges := make([][]int,n)
    Flag := make([]int,n)
    for i:=0;i<len(graph);i++{
        edges[graph[i][0]] = append(edges[graph[i][0]],graph[i][1])
    }
    queue := make([]int,0)
    queue = append(queue,start)
    for len(queue) > 0{
        tmp := queue[0]
        queue = queue[1:]
        Flag[tmp] = 1
        if tmp == target{
            return true
        }
        for i:=0;i<len(edges[tmp]);i++{
            if Flag[edges[tmp][i]] == 0{
                queue = append(queue,edges[tmp][i])
            }
        }
    }
    return false

}