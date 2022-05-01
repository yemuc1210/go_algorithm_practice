package main
import (
	"fmt"
)
// 子岛屿：grid2一个岛屿被grid1的岛屿完全包含
// 求grid2中子岛屿的数目
// 使用并查集？  grid[i][j] = 所属岛屿的编号？ 编号  dfs的做法了吧

// 1 dfs
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	// 遍历grid2 判断相同位置grid1的情况
	if len(grid1)==0 || len(grid2)==0 {return 0}
	const (
		LAND = 1
		WATER = 0
	)
	m,n := len(grid1),len(grid1[0])  // m*n
	var dfs func(grid [][]int, i,j int)
	dfs = func(grid [][]int,i,j int) {
		// 边界
		if i<0 || j<0 || i>=m || j>=n {
			return 
		}
		if grid[i][j] == WATER {return }
		grid[i][j] = WATER  // 打个标记
		dfs(grid, i-1, j) // 上
		dfs(grid, i+1, j) // 下
		dfs(grid, i, j-1) // 左
		dfs(grid, i, j+1) // 右
	}
	
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid1[i][j]==WATER && grid2[i][j] == LAND {
				dfs(grid2,i,j)   // 淹没非子岛屿的岛屿
			}
		}
	}
	// 统计剩余岛屿的数量,
	res := 0 
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid2[i][j] == LAND {
				res ++
				dfs(grid2,i,j)
			}
		}
	}
	return res
}

func main() {
	grid1 := [][]int{
		{1,1,1,0,0},{0,1,1,1,1},{0,0,0,0,0},{1,0,0,0,0},{1,1,0,1,1},
	}
	grid2 := [][]int{
		{1,1,1,0,0},{0,0,1,1,1},{0,1,0,0,0},{1,0,1,1,0},{0,1,0,1,0},
	}
	res := countSubIslands(grid1,grid2)
	fmt.Println(res)
}