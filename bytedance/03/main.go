package main

import "fmt"

func main() {
	var MAX int
	var L int
	var W [][]int
	fmt.Scanln(&MAX)
	fmt.Scanln((&L))
	W = make([][]int, L)
	for i:= range W {
		W[i] = make([]int, L)
	}
	// 输入成本
	for i:=0;i<L;i++{
		for j:=0;j<L;j++{
			fmt.Scanf("%d%c",&W[i][j])
		}
	}
	// 测试   -ok
	// fmt.Println(W)

	// 计算
	res := solve(MAX, W)
	fmt.Println(res)
}
// 输出最长路线的长度
func solve(totalCost int, costMatrix [][]int) int {
	// 每个放假只能铺一次，且不能交叉
	// 最终从外围放假出去
	// totalCost 花完，且路径最长
	maxLength := 0
	n := len(costMatrix)
	visited := make([][]bool,n)
	for i:=range visited {visited[i] = make([]bool, n)}

	dirs := [][]int{
		{-1,0},{1,0},{0,-1},{0,1},
	}

	var dfs func(leftCost int, pathLength int, i,j int)
	dfs = func(leftCost, pathLength int, i,j int) {
		if leftCost == 0 {
			// 预算画完，需要判断能否出去
			if i==n-1 && j==n-1 {
				// 当前位置在边界，可以出去，合法
				if pathLength >= maxLength {
					maxLength = pathLength
				}
			}
			fmt.Println("pathLength:",pathLength)
			return 
		}
		// 当前层逻辑 可以选择四个方向 上下左右
		// 先访问i,j
		visited[i][j] = true
		for _,dir := range dirs {
			x,y := dir[0],dir[1]
			nextI,nextJ := i+x,j+y
			// 判断能否进入下一层遍历
			if nextI>=0 && nextI<n && nextJ>=0 &&nextJ<n && !visited[nextI][nextJ] {
				dfs(leftCost-costMatrix[i][j], pathLength+1, nextI,nextJ)
			}
		}
		// 状态回溯 重置
		visited[i][j] = false
	}
	dfs(totalCost-costMatrix[0][0],1,0,0)
	// fmt.Println("maxLength:",maxLength)
	return maxLength
}
