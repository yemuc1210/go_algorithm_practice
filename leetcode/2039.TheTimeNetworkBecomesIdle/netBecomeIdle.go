package lt2039

func networkBecomesIdle(edges [][]int, patience []int) int {
	var res int
	// n个服务器，编号0-n-1  边edges[i]= [ui,vi]表示服务器之间存在通路
	// 一秒内可以传输任意数目的信息
	// patience，长度n，下标从0开始
	// 编号0为主服务器，其他为数据服务器。
	// 信息在服务器之间以最优线路传输，数据服务器向主服务器发送信息并以最少的时间到达主服务器
	// 数据服务器i 存在patience[i]的重发时间
	// 主服务器处理所有新到达信息，并立即反方向发送回复信息
	// 问题：求任意数据服务器到达主服务器的最短路径中 最长的路径长度
	// 求每个数据服务器到达主服务器的最短路径   单源最短路径问题-Dijstra？   BFS

	// 1. 建表
	n := len(patience)
	grapg := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		grapg[x] = append(grapg[x], y)
		grapg[y] = append(grapg[y], x)
	}
	// 2. BFS 求得每个节点到主服务器的最短路径 dist[]  dist[i]:表示节点i到0的最短路径
	visited := make([]bool, n)
	visited[0] = true
	queue := []int{0} // 先访问主服务器
	for dist := 1; len(queue) > 0; dist++ {
		// 遍历当前队列中所有节点
		tmp := queue
		queue = nil
		for _, x := range tmp {
			// 访问节点x的相邻节点   图以邻接矩阵形式存储
			for _, v := range grapg[x] {
				if visited[v] {
					continue
				}
				visited[v] = true
				queue = append(queue, v)
				res = max(res, (dist*2-1)/patience[v]*patience[v]+dist*2+1)
			}
		}
	}
	// 3. 考虑重发时间+回复时间  di<=t=patience[i] 数据服务器不会重发，该节点停止时间为dist
	// di>t  数据服务器会发生重发，最后一个数据包的发送时间为  (dist[i] - 1) /t*t  停止时间为(dist[i]-1)/t*t+dist[i]

	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
