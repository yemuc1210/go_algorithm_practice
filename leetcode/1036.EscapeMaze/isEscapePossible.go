package lt1036
/*
1036. 逃离大迷宫

在一个 106 x 106 的网格中，每个网格上方格的坐标为 (x, y) 。

现在从源方格 source = [sx, sy] 开始出发，意图赶往目标方格 target = [tx, ty] 。
数组 blocked 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。

每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格 不 在给出的封锁列表 blocked 上。同时，不允许走出网格。

只有在可以通过一系列的移动从源方格 source 到达目标方格 target 时才返回 true。否则，返回 false。
*/

//dfs bfs  不过考虑问题规模大，要优化
//改搜目标方格是否被包围  最优的包围方式是斜着围成等腰三角形
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isEscapePossible(block [][]int, source, target []int) bool {
    const (
        blocked = -1 // 在包围圈中
        valid   = 0  // 不在包围圈中
        found   = 1  // 无论在不在包围圈中，但在 n(n-1)/2 步搜索的过程中经过了 target

        boundary int = 1e6
    )

    n := len(block)
    if n < 2 {
        return true
    }

    blockSet := map[pair]bool{}
    for _, b := range block {
        blockSet[pair{b[0], b[1]}] = true
    }

    check := func(start, finish []int) int {
        sx, sy := start[0], start[1]
        fx, fy := finish[0], finish[1]
        countdown := n * (n - 1) / 2

        q := []pair{{sx, sy}}
        vis := map[pair]bool{{sx, sy}: true}
        for len(q) > 0 && countdown > 0 {
            p := q[0]
            q = q[1:]
            for _, d := range dirs {
                x, y := p.x+d.x, p.y+d.y
                np := pair{x, y}
                if 0 <= x && x < boundary && 0 <= y && y < boundary && !blockSet[np] && !vis[np] {
                    if x == fx && y == fy {
                        return found
                    }
                    countdown--
                    vis[np] = true
                    q = append(q, np)
                }
            }
        }

        if countdown > 0 {
            return blocked
        }
        return valid
    }

    res := check(source, target)
    return res == found || res == valid && check(target, source) != blocked
}