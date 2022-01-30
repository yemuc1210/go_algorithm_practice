package NC109

/**
 * 判断岛屿数量 lt200
 * @param grid char字符型二维数组 
 * @return int整型
*/
func solve( grid [][]byte ) int {
    // 遍历 bfs  dfs  类似连通子图
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] == '1'{
				dfs(grid,i,j)
				cnt ++
			}
		}
	}
	return cnt
}

func dfs(grid [][]byte, i, j int) {
	if (i < 0 || j < 0) || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
