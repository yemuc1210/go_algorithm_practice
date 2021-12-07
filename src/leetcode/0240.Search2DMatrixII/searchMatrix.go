package lt240
/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
 

Z字查找
*/

func searchMatrix(matrix [][]int, target int) bool {//右上角 (0, n-1)(0,n−1) 进行搜索
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row++
		} else {
			col--
		}
	}
	return false
}