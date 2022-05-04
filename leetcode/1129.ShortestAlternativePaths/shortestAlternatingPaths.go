package main
// // 1129. 颜色交替的最短路径
/**
  创建一个n*n的邻接矩阵来表示图，g[i][j]表示i节点到j节点的路线的颜色
  1表示红色边，-1表示蓝色边，0表示同时存在两种颜色的边，2表示i和j没有连通
*/
func buildGraph(n int, redEdges, blueEdges [][]int) [][]int {
	g := make([][]int, n)
	// 初始化所有边都是2
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = 2
		}
	}
    // 遍历redEdges, 在g中标记所有的红色边
	for i := range redEdges {
		x, y := redEdges[i][0], redEdges[i][1]
		g[x][y] = 1
	}
    // 遍历blueEdges, 如果两个节点已经以红色边相连，表示这两个节点存在两种颜色的边
	for i := range blueEdges {
		x, y := blueEdges[i][0], blueEdges[i][1]
		if g[x][y] == 1 {
			g[x][y] = 0
		} else {
			g[x][y] = -1
		}

	}
	return g
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	g := buildGraph(n, redEdges, blueEdges)
	// visited表示以某种颜色的路径到达某个节点的路径是否已经处理过
    // 以不同颜色的路径到达同一节点是两种不同情况，需要保存两个状态
    // 我们以一个拼接的string来描述这两种状态
	visited := make(map[string]bool)
    // queue中的元素是一个int数组，每个int数组中只有两个元素
    // 其中[0]表示进入[1]节点的路径的颜色
	queue := make([][]int, 0)
    // queue中首先加入的两个元素，分别表示进入0节点的两条路径
    // 因为红蓝交叉的路径可以以两种颜色开始，可以等价为从dummy节点进入0节点的路径颜色可以是红色和蓝色
    // append的这两个数组就解释为：以红色路径（1）进入0节点，以蓝色路径（-1）进入0节点
	queue = append(queue, []int{0, 1}, []int{0, -1})
    // res[i]表示0到i的红蓝交叉的最短路径
	res := make([]int, n)
    // 初始化res[i]为-1表示不可到达
	for i := 0; i < n; i++ {
		res[i] = -1
	}
    // length表示路径的长度，因为采用的宽搜，所有每有一波元素（而不是一个，这是bfs基本思想）进入queue，就表示路径多了一步
	length := 0
	for len(queue) > 0 {
        // 每次进入for表示一层bfs结束，路径长度++
		length++
        // size表示这一层bfs的节点个数
		size := len(queue)
		for i := 0; i < size; i++ {
            // 记住我们queue中保存的数组的意义
			node := queue[0][0]
			color := queue[0][1]
            // 我们定义的红色和蓝色对应的数字是相反的，可以很容易的拿到下一段路径的颜色
			oppoColor := -color
			queue = queue[1:]
            // 这里从0开始遍历所有节点，就是要找出我们创建的图g中，有没有符合要求的下一条边
            // 以第一次进入循环的状态为例，此时queue中弹出的数组元素是{0,1}表示红边进入0节点
            // 那么就要在g中找有没有以蓝边从0节点到另一个节点的路径，即找g[0][j]==-1||g[0][j]==0是否存在
			for j := 0; j < n; j++ {
				if g[node][j] == oppoColor || g[node][j] == 0 {
                    // 拼接visited的key，有两个影响因素：进入的节点和进入时边的颜色
					key := fmt.Sprintf("%d%d", j, oppoColor)
					if visited[key] {
						continue
					}
					visited[key] = true
                    // 发现了符合条件的边，将节点和边的颜色入队列
					queue = append(queue, []int{j, oppoColor})
                    // 发现了到达j节点的一条路径，可以更新res数组了
                    // 如果res[j]还是初始状态，直接更新为length，否则就要动态维护一个最小值
					if res[j] == -1 {
						res[j] = length
					} else {
						res[j] = min(res[j], length)
					}
				}
			}
		}
	}
    // 0节点是特殊情况，最后特殊处理一下
	res[0] = 0
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

