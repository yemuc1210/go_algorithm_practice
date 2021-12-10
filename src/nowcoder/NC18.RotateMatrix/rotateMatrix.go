package NC18

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 顺时针旋转数组   中等   面试题01.07 
 * @param mat int整型二维数组 
 * @param n int整型 
 * @return int整型二维数组
*/
func rotateMatrix( mat [][]int ,  n int ) [][]int {
    // 矩阵顺时针旋转90°
	//思路先对角线变换  然后竖直轴对称翻转
	// row := len(mat)
	if n <= 0 {
		return mat
	}
	row,col := n,n
	for i:=0;i<row;i++{
		for j:=i+1;j<col;j++{
			mat[i][j],mat[j][i] = mat[j][i],mat[i][j]
		}
	}

	//竖直轴对称翻转
	halfCol := col/2
	for i:=0;i<row;i++{
		for j:=0;j<halfCol;j++{
			mat[i][j],mat[i][col-j-1] = mat[i][col-j-1],mat[i][j]
		}
	}
	return mat
}