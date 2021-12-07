package lt598

import "math"

/*
598. 范围求和 II
给定一个初始元素全部为 0，大小为 m*n 的矩阵 M 以及在 M 上的一系列更新操作。

操作用二维数组表示，其中的每个操作用一个含有两个正整数 a 和 b 的数组表示，
含义是将所有符合 0 <= i < a 以及 0 <= j < b 的元素 M[i][j] 的值都增加 1。

在执行给定的一系列操作后，你需要返回矩阵中含有最大整数的元素个数。
*/
func maxCount(m int, n int, ops [][]int) int {
	if m<1 || n<1 {  //不存在二维数组
		return 0
	}
	if len(ops) == 0 {
		return m*n
	}
	//找到行最小，列最小
	minCol,minRow := math.MaxInt32,math.MaxInt32
	for i:=0;i<len(ops);i++{
		if ops[i][0] < minCol {
			minCol = ops[i][0]
		}
		if ops[i][1] < minRow {
			minRow = ops[i][1]
		}
	}
	return minCol*minRow
}