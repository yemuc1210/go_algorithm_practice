package offer29

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
*/

//按层模拟  左上角-右下角规定好  然后第二层就是左上角坐标+1  右下角坐标-1
//8ms 98.85%   6.5MB  36.02%
func spiralOrder(matrix [][]int) []int {
	m := len(matrix) //行
	if m==0{
		return nil
	}
	n:=len(matrix[0]) //列
	if n==0{
		return nil
	}
	//剩余区域的上左右下
	top,left,bottom,right := 0,0,m-1,n-1

	cnt,sum:=0,m*n
	res:=[]int{}
	for cnt<sum{
		i,j := top,left
		for j<=right && cnt<sum {
			res = append(res, matrix[i][j])
			cnt++
			j++
		}
		i,j = top+1,right
		for i<=bottom && cnt<sum{
			res = append(res, matrix[i][j])
			i++
			cnt++
		}
		i,j =bottom,right-1
		for j>=left && cnt<sum{
			res = append(res, matrix[i][j])
			j--
			cnt++
		}
		i,j =bottom-1,left
		for i>top && cnt<sum{
			res = append(res, matrix[i][j])
			i--
			cnt++
		}   //完成一圈

		top,left,bottom,right = top+1,left+1,bottom-1,right-1

	}

	return res

}