package tencent50
/*中等
54. 螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/
/*
烦 
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0])==0{
		return []int{}
	}
	res := []int{}
	
	//行列
	m,n := len(matrix),len(matrix[0])
	left,right,top,bottom := 0,n-1,0,m-1
	/*
	终止条件：
		如果行是偶数 left right相邻   
	*/
	for {  //终止条件如何写 待定
		//访问上面一行  即top所在的那行
		for i:=left;i<=right;i++{
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom{
			break
		}
		//访问最后一列   matrix[][right]
		for j:=top;j<=bottom;j++{
			res = append(res, matrix[j][right])
		}
		right--
		if right<left{
			break
		}
		//访问最底一行  matrix[bottom][]
		for i:=right;i>=left;i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if bottom < top{
			break
		}
		//访问最左一列    matrix[][left]
		for j:=bottom;j>=top;j--{
			res = append(res, matrix[j][left])
		}
		left++
		if left>right{
			break
		}
		//终止条件 left==right
	}
	return res
}