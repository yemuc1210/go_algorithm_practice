package leetcode


/*
我们有一系列公交路线。每一条路线 routes[i] 上都有一辆公交车在上面循环行驶。
例如，有一条路线 routes[0] = [1, 5, 7]，表示第一辆 (下标为0) 
公交车会一直按照 1->5->7->1->5->7->1->... 的车站路线行驶。
假设我们从 S 车站开始（初始时不在公交车上），要去往 T 站。
 期间仅可乘坐公交车，求出最少乘坐的公交车数量。
 返回 -1 表示不可能到达终点车站。
*/


/*
思路：
转换成图论的问题，将每个站台看成顶点，公交路径看成每个顶点的边。
同一个公交的边染色相同。题目即可转化为从顶点 S 到顶点 T 需要经过最少多少条不同的染色边。
用 BFS 即可轻松解决。从起点 S 开始，不断的扩展它能到达的站点。
用 visited 数组防止放入已经可达的站点引起的环。
用 map 存储站点和公交车的映射关系(即某个站点可以由哪些公交车到达)，
BFS 的过程中可以用这个映射关系，拿到公交车的其他站点信息，从而扩张队列里面的可达站点。
一旦扩展出现了终点 T，就可以返回结果了。
*/

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source==target{
		return 0
	}
	visited,queue,res,vertexMap := make([]bool,len(routes)),[]int{},0,map[int][]int{}
	//map  key是顶点，value是公交车数组，表示这些公交车可以到这个站点
	for i:=0;i<len(routes);i++{
		for _,v := range routes[i]{   //遍历每条路线，判断能否到达该站点
			tmp:=vertexMap[v]   //等到达站点v的公交车数组
			tmp=append(tmp, i) //route[i]表示公交车i的路线
			vertexMap[v] = tmp
		}
	}
	queue = append(queue, source)

	for len(queue) > 0 {
		res++
		qlen:=len(queue)
		for i:=0;i<qlen;i++{
			vertex := queue[0]
			queue = queue[1:]//出队
			for _,bus := range vertexMap[vertex]{
				if visited[bus]{
					continue
				}
				visited[bus] = true
				for _,v := range routes[bus]{
					if v==target{
						return res
					}
					queue = append(queue, v)
				}
			}
		}
	}
	return -1
}