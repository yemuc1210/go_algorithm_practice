package main

// 单元格比上下左右高，则雨水会向下流
// dfs搜索，从矩阵的边界逆向搜索：边界最低，向上寻转高点
// 由于矩阵的左边界和上边界是太平洋，矩阵的右边界和下边界是大西洋，
// 因此从矩阵的左边界和上边界开始反向搜索即可找到雨水流向太平洋的单元格，
// 从矩阵的右边界和下边界开始反向搜索即可找到雨水流向大西洋的单元格。

func pacificAtlantic(heights [][]int) [][]int {
	// 找到每一行的最高值？
	m,n := len(heights),len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool,m)
	for i:=range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}
	// 
	var dirs = [][]int {
		{-1,0},{1,0},{0,-1},{0,1},
	}
	var dfs func(i,j int, ocean [][]bool)
	dfs = func(i, j int, ocean [][]bool) {
		if ocean[i][j] {
			return 
		} 
		// visit
		ocean[i][j] = true
		// 四个方向
		for _,dir := range dirs{
			nx,ny := i+dir[0],j+dir[1]
			// 判断高度
			// if heights[i][j] <= heights[nx][ny] {
			// 	dfs(nx,ny, ocean)
			// }
			if 0 <= nx && nx < m && 0 <= ny && ny < n && heights[nx][ny] >= heights[i][j] {
                dfs(nx, ny, ocean)
            }
		}
	}
	// 从四个边界
	for i := 0; i < m; i++ {
        dfs(i, 0, pacific)
    }
    for j := 1; j < n; j++ {
        dfs(0, j, pacific)
    }
    for i := 0; i < m; i++ {
        dfs(i, n-1, atlantic)
    }
    for j := 0; j < n-1; j++ {
        dfs(m-1, j, atlantic)
    }
	var res [][]int
	// 遍历
	for i,row := range pacific {
		for j,ok := range row {
			if ok&&atlantic[i][j] {
				// all ok
				res = append(res, []int{i,j})
			}
		}
	}
	return res
}