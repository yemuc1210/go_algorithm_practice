package NC86

/**
 * 矩阵元素查找
 * @param mat int整型二维数组 
 * @param n int整型 
 * @param m int整型 
 * @param x int整型 
 * @return int整型一维数组  坐标（x,y)
*/
func findElement( mat [][]int ,  n int ,  m int ,  x int ) []int {
    // 行列皆有序  二分    n行m列
	res := []int{}

	i,j := 0,m-1
	for i<n && j>=0 {
		if mat[i][j] > x {
			//那么元素在左侧
			j--
		}else if mat[i][j] < x {
			//比当前右上角还大，只能在其下面
			i ++
		}else {
			res = append(res, i)
			res = append(res, j)
			break
		}
	}
	return res
}