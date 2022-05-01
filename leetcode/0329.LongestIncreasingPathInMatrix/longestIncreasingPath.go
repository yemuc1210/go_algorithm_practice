package lt329
/*lt329 图
剑指 Offer II 112. 最长递增路径
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
*/
import (
	"math"
)

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func longestIncreasingPath(matrix [][]int) int {
	cache, res := make([][]int, len(matrix)), 0
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(matrix[0]))
	}
	for i, v := range matrix {
		for j := range v {
			searchPath(matrix, cache, math.MinInt64, i, j)
			res = max(res, cache[i][j])
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isInIntBoard(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchPath(board, cache [][]int, lastNum, x, y int) int {
	if board[x][y] <= lastNum {
		return 0
	}
	if cache[x][y] > 0 {
		return cache[x][y]
	}
	count := 1
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInIntBoard(board, nx, ny) {
			count = max(count, searchPath(board, cache, board[x][y], nx, ny)+1)
		}
	}
	cache[x][y] = count
	return count
}
