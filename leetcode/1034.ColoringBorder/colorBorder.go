package lt1034
/*
1034. 边界着色
给你一个大小为 m x n 的整数矩阵 grid ，表示一个网格。另给你三个整数 row、col 和 color 。
网格中的每个值表示该位置处的网格块的颜色。

当两个网格块的颜色相同，而且在四个方向中任意一个方向上相邻时，它们属于同一 连通分量 。

连通分量的边界 是指连通分量中的所有与不在分量中的网格块相邻（四个方向上）的所有网格块，
	或者在网格的边界上（第一行/列或最后一行/列）的所有网格块。

请你使用指定颜色 color 为所有包含网格块 grid[row][col] 的 连通分量的边界 进行着色，并返回最终的网格 grid 。

 dfs
*/
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder(grid [][]int, row, col, color int) [][]int {
    m, n := len(grid), len(grid[0])
    type point struct{ x, y int }
    borders := []point{}
    originalColor := grid[row][col]
    vis := make([][]bool, m)
    for i := range vis {
        vis[i] = make([]bool, n)
    }

    var dfs func(int, int)
    dfs = func(x, y int) {
        vis[x][y] = true
        isBorder := false
        for _, dir := range dirs {
            nx, ny := x+dir.x, y+dir.y
            if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
                isBorder = true
            } else if !vis[nx][ny] {
                vis[nx][ny] = true
                dfs(nx, ny)
            }
        }
        if isBorder {
            borders = append(borders, point{x, y})
        }
    }
    dfs(row, col)

    for _, p := range borders {
        grid[p.x][p.y] = color
    }
    return grid
}

