package main

import "sort"
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath1(heights [][]int) int {
    n, m := len(heights), len(heights[0])
    return sort.Search(1e6, func(maxHeightDiff int) bool {
        vis := make([][]bool, n)
        for i := range vis {
            vis[i] = make([]bool, m)
        }
        vis[0][0] = true
        queue := []pair{{}}
        for len(queue) > 0 {
            p := queue[0]
            queue = queue[1:]
            if p.x == n-1 && p.y == m-1 {
                return true
            }
            for _, d := range dirs {
                x, y := p.x+d.x, p.y+d.y
                if 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && abs(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
                    vis[x][y] = true
                    queue = append(queue, pair{x, y})
                }
            }
        }
        return false
    })
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}