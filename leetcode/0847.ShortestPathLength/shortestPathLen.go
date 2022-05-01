package lt847
/*
存在一个由 n 个节点组成的无向连通图，图中的节点按从 0 到 n - 1 编号。

给你一个数组 graph 表示这个图。其中，graph[i] 是一个列表，由所有与节点 i 直接相连的节点组成。

返回能够访问所有节点的最短路径的长度。你可以在任一节点开始和停止，也可以多次重访节点，并且可以重用边。


从某个点开始访问所有节点   迪杰斯特拉   BFS
(u,mask,distLen)
从任意节点开始（i,2^i,0）  如果mask包含n个1,即mask=2^n-1 则成功，返回distLen

为了保障节点u不被多次访问，使用数组或者哈希表记录(u,mask)是否被搜索

*/
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type triTuple struct{
		u,mask,distLen int
	}
	q := []triTuple{}  //BFS遍历 队列
	seen := make([][]bool,n)
	for i:=range seen{
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, triTuple{i,1<<i,0})
	}
	for {
		cur := q[0]
		q = q[1:]   //出队
		if cur.mask == 1<<n-1 {
			return cur.distLen
		}
		//搜索相邻
		for _,v := range graph[cur.u] {
			maskV := cur.mask | 1<<v
			if !seen[v][maskV] {
				q = append(q, triTuple{v,maskV,cur.distLen+1})
				seen[v][maskV] = true
			}
		}
	}
}