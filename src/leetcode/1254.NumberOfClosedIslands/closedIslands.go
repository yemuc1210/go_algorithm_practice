package main
var coordinate = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func closedIsland(grid [][]int) int {
	flag := make([][]bool, len(grid))
	for index := range flag {
		flag[index] = make([]bool, len(grid[0]))
	}
	//先遍历行边
	for i := 0; i < len(grid); i++ {
		dfs(flag, i, 0, grid)
		dfs(flag, i, len(grid[0]) - 1, grid)
	}
	//遍历列边
	for j := 0; j < len(grid[0]); j++ {
		dfs(flag, 0, j, grid)
		dfs(flag, len(grid) - 1, j, grid)
	}
	
	var ans int = 0
	for i := 1; i < len(grid) - 1; i++ {
		for j := 1; j < len(grid[0]) - 1; j++ {
			if grid[i][j] == 0 {
				ans++
				dfs(flag, i, j, grid)
			}
		}
	}
	return ans
}

func dfs(flag [][]bool, row int, col int, grid [][]int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || flag[row][col] {
		return
	}
	flag[row][col] = true
	if grid[row][col] == 1 {
		return
	}
	grid[row][col] = 1

	for _, value := range coordinate {
		dfs(flag, row+value[0], col+value[1], grid)
	}
}

 