package main
// 1926. 迷宫中离入口最近的出口
// 空格. 墙+
// 给定entrance 入口
// 上下左右移动
// 出口:边界上的空格
// 返回最短路径的步数
type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit(g [][]byte, entrance []int) int {
    n, m := len(g), len(g[0])
    s := pair{entrance[0], entrance[1]}
    g[s.x][s.y] = 0
    q := []pair{s}
    for ans := 1; len(q) > 0; ans++ {
        tmp := q
        q = nil
        for _, p := range tmp {
            for _, d := range dir4 {
                if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' {
                    if x == 0 || y == 0 || x == n-1 || y == m-1 {
                        return ans
                    }
                    g[x][y] = 0
                    q = append(q, pair{x, y})
                }
            }
        }
    }
    return -1
}