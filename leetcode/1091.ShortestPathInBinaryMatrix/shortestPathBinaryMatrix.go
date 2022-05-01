package lt1091
/*
20天算法计划    BFS-DFS

给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。

二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：

路径途经的所有单元格都的值都是 0 。
路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
畅通路径的长度 是该路径途经的单元格总数。


*/


func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid)==0 || len(grid[0])==0 {
		return -1
	}
	if grid[0][0] == 1{
		return -1   //起点阻塞
	}
	res := 0 
	m,n := len(grid),len(grid[0])
	direction := [][]int{
		{-1,0},{+1,0},{0,-1},{0,1},{-1,-1},{-1,1},{1,1},{1,-1},
	}
	queue := [][]int{}  //存放位置坐标的队列

	queue = append(queue,   []int{0,0})

	res += 1

	for len(queue) != 0 {
		size := len(queue)
		for ;size>0;size--{
			cur := queue[0]
			queue = queue[1:]
			x,y := cur[0],cur[1]

			if x==m-1 && y==n-1 {
				//到达终点
				return res
			}
			//八个方向判断
			for _,dir:= range direction {
				x1,y1 := x+dir[0], y+dir[1]

				//过滤无效方向
				if x1<0 || x1>=m || y1<0 || y1>=n || grid[x1][y1] == 1{
					continue
				}	
				//否则为有效方向
				queue = append(queue, []int{x1,y1})
				grid[x1][y1] = 1   //访问过的标记
			}
		}
		res++   //遍历完这一层所有节点，步数++
	}



	return -1
}

