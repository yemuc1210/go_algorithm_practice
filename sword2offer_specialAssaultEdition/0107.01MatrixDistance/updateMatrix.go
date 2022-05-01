package offerS107

/*剑指 Offer II 107. 矩阵中的距离
给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，
其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。
*/

/*
BFS搜索四个方向
*/
func updateMatrix(mat [][]int) [][]int{
	n,m := len(mat),len(mat[0])
	queue := make([][]int,0)
	for i:=0;i<n;i++{// 0 存入队列
		for j:=0;j<m;j++{
			if mat[i][j] == 0{
				point := []int{i,j}
				queue = append(queue, point)
			}else{
				mat[i][j] = -1
			}
		}
	}
	//定义四个方向
	direction := [][]int{{0,1},{0,-1},{1,0},{-1,0}}

	//BFS操作
	for len(queue) > 0{ //!empty
		point := queue[0]
		queue = queue[1:]  //取队头
		//遍历四个方向
		for _,v := range direction {
			x,y := point[0]+v[0], point[1]+v[1]

			if x>=0 && x<n && y>=0 && y<m && mat[x][y]==-1{//未访问过，原来是1的
				mat[x][y] = mat[point[0]][point[1]] + 1 //一步遍历得到
				queue = append(queue, []int{x,y})
			}
		}
	}
	return mat
}