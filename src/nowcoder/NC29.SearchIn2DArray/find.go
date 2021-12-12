package NC29

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 二维数组中的查找    lt240
 * @param target int整型 
 * @param array int整型二维数组 
 * @return bool布尔型
*/
func Find( target int ,  matrix [][]int ) bool {
    // 二分
	if len(matrix) == 0 {
		return false
	}
	row,col := 0,len(matrix[0])-1   //row 行  col列

	//与右对比
	for col>=0 && row<=len(matrix)-1 {
		if target == matrix[row][col]{
			return true
		}else if target > matrix[row][col] {
			row++
		}else{
			col --
		}
	}
	return false
}