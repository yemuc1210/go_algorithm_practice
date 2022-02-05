package lt1219

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getMaximumGold(grid [][]int) (ans int) {
    var dfs func(x, y, gold int)
    dfs = func(x, y, gold int) {
        gold += grid[x][y]
        if gold > ans {
            ans = gold
        }

        rec := grid[x][y]
        grid[x][y] = 0
        for _, d := range dirs {
            nx, ny := x+d.x, y+d.y
            if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[nx]) && grid[nx][ny] > 0 {
                dfs(nx, ny, gold)
            }
        }
        grid[x][y] = rec
    }

    for i, row := range grid {
        for j, gold := range row {
            if gold > 0 {
                dfs(i, j, 0)
            }
        }
    }
    return
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/path-with-maximum-gold/solution/huang-jin-kuang-gong-by-leetcode-solutio-f9gg/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。