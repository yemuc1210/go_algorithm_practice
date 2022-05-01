package main

import "fmt"

// 找出海洋格（0），离它最近的陆地单元（1)的距离是最大的
func maxDistance(grid [][]int) int {
	n := len(grid) // n*n
	// 遍历海洋格
	var dfs func(i, j int, srci, srcj int, minDis *int, visitedFlag int)
	dfs = func(i, j int, srci, srcj int, minDis *int, visitedFlag int) {
		if i < 0 || i >= n || j < 0 || j >= n || visitedFlag == grid[i][j] {
			return
		}
		if grid[i][j] == 1 {
			*minDis = min(*minDis, getDistance(i, srci, j, srcj))
			return
		}
		grid[i][j] = visitedFlag
		// 上下左右
		dfs(i+1, j, srci, srcj, minDis, visitedFlag)
		dfs(i-1, j, srci, srcj, minDis, visitedFlag)
		dfs(i, j+1, srci, srcj, minDis, visitedFlag)
		dfs(i, j-1, srci, srcj, minDis, visitedFlag)
	}
	maxDis := 0
	visitedFlag := 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				var minDis int
				dfs(i, j, i, j, &minDis, visitedFlag)
				maxDis = max(maxDis, minDis)
				visitedFlag++
			}
		}
	}
	return maxDis
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getDistance(x0, x1, y0, y1 int) int {
	return abs(x0, x1) + abs(y0, y1)
}

// bfs 解法
// 首先 陆地入队
// 然后从各个陆地 一圈一圈的遍历海洋  最后遍历到的就是最远的
type Point struct {
    X int
    Y int
}

func maxDistance1(grid [][]int) int {
    var queue []*Point
    for i:=0; i < len(grid); i++ {
        for j:=0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                queue = append(queue, &Point{i, j})
            }
        }
    }
    if len(queue) == 0 || len(queue) == len(grid)*len(grid[0]) {
        return -1
    }

    ans := 0
    d := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    for len(queue) > 0 {
        ans++
        tmpQueu := queue
        queue = nil
        for len(tmpQueu) > 0 {
            p := tmpQueu[0]
            tmpQueu = tmpQueu[1:]
            // 以p为原点，检查4个方向
            for i:=0; i < 4; i++ {
                x := p.X + d[i][0]
                y := p.Y + d[i][1]
                if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != 0 {
                    continue
                }
                queue = append(queue, &Point{x, y})
                grid[x][y] = 2 // 代表以及被遍历过了
            }
        }
    }

    return ans-1
}


func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func main() {
	fmt.Println("vim-go")
	grid := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1}}
	maxDis := maxDistance1(grid)
	fmt.Println(maxDis)
}
