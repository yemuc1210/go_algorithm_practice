package lt74

import "sort"

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。


基本上就是个一维升序的感觉，左上角最小，右下角最大
二分

一次二分

*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0 {
		return false
	}
	
	m,n := len(matrix),len(matrix[0])
	
	i := sort.Search(m*n, func(i int)bool{
		return matrix[i/n][i%n]>=target
	})
	return i<m*n  && matrix[i/n][i%n]==target
}