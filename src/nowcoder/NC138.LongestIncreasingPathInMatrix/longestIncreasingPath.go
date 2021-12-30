package NC138

import "math"

//lt329

/*
描述
给定一个 n 行 m 列矩阵 matrix ，矩阵内所有数均为非负整数。 你需要在矩阵中找到一条最长路径，使这条路径上的元素是递增的。并输出这条最长路径的长度。
这个路径必须满足以下条件：

1. 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外。
2. 你不能走重复的单元格。即每个格子最多只能走一次。
*/

/*
dfs
不能走重复的，用visited记录
一个全局变量记录当前最大值
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 递增路径的最大长度
 * @param matrix int整型二维数组 描述矩阵的每个数
 * @return int整型
 */
var dir = [][]int{
	{-1,0},
	{0,-1},
	{0,1},
	{1,0},
}
func solve( matrix [][]int ) int {
    // write code here
	// visited := make([][]bool, len(matrix))
	//记忆化搜索 使用cache矩阵作为缓存矩阵 记录已经计算过的
	cache := make([][]int, len(matrix))
	for i:=range cache{
		cache[i] = make([]int, len(matrix[i]))
	}
	res := 0

	//遍历
	for i,row := range matrix{
		for j := range row {
			//从当前位置（i，j）开始搜
			searchPath(matrix, cache, math.MinInt64, i, j)
			res = Max(res, cache[i][j])
		}
	}
	return res
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
		if isValidIndex(board, nx, ny) {
			count = Max(count, searchPath(board, cache, board[x][y], nx, ny)+1)
		}
	}
	cache[x][y] = count
	return count
}

func isValidIndex(board [][]int, x,y int) bool {
	return x>=0 && x<len(board) && y>=0 && y<len(board[0])
}

func Max(a,b int) int {
	if a>b {
		return a
	}
	return b
}