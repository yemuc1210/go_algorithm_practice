package main

import (
	"fmt"
)

// 最大人工岛
// 最多只能将一个0变成1
// 返回grid中最大的岛屿面积  岛屿：1形成的
// 最大的岛屿面积 要么来自于原来单个最大的岛屿，要么是两个岛屿可以合并（通过一个新增的1）
// dfs
// 并查集

var dr = []int{-1, 0, 1, 0}
var dc = []int{0, -1, 0, 1}
var N int
var grid [][]int

func largestIsland(g [][]int) int {
	// dfs 每个连通块中节点的值变为连通块的大小
	grid = g
	N = len(grid)
	index := 2
	area := make([]int, N*N+2)
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if grid[r][c] == 1 {
				area[index] = dfs(r, c, index)
				index++ // index 是不同连通块的编号

			}
		}
	}
	//
	res := 0

	for _, x := range area {
		res = max(res, x) // area[i] i所在联通分块的面积
	}
	//
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if grid[r][c] == 0 {
				seen := make(map[int]struct{}) // 空struct不占空间
				for _, move := range neighbors(r, c) {
					if grid[move/N][move%N] > 1 {
						seen[grid[move/N][move%N]] = struct{}{}
					}
				}
				bns := 1
				for i := range seen {
					bns += area[i]
				}
				res = max(res, bns)
			}
		}
	}

	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(r, c, index int) int {
	ans := 1
	grid[r][c] = index
	for _, move := range neighbors(r, c) {
		if grid[move/N][move%N] == 1 {
			grid[move/N][move%N] = index
			ans += dfs(move/N, move%N, index)
		}
	}
	return ans
}
func neighbors(r, c int) []int {
	l := []int{}
	for k := 0; k < 4; k++ {
		nr := r + dr[k]
		nc := c + dc[k]
		if 0 <= nr && nr < N && 0 <= nc && nc < N {
			l = append(l, nr*N+nc)
		}
	}
	return l
}

func main() {
	grid := [][]int{
		{1, 0}, {0, 1},
	}
	res := largestIsland(grid)
	fmt.Println(res)
}
