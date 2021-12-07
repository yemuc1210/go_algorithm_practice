package NC77

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *	offer21 不同之处在于本题要求相对位置不变
 * 调整数组顺序使奇数位于偶数前面
 * @param array int整型一维数组 
 * @return int整型一维数组
*/
func reOrderArray( nums []int ) []int {
    // 双指针
	idx := 0
	res := make([]int, len(nums))
	for _,v:=range nums {
		if v%2 != 0 {
			//奇数
			res[idx] = v
			idx ++
		}
	}
	//处理偶数
	for _,v := range nums{
		if v%2 == 0 {
			res[idx] = v
			idx ++
		}
	}
	return res
}