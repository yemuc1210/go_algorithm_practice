package NC110

/**
 * 旋转数组
 * @param n int整型 数组长度
 * @param m int整型 右移距离
 * @param a int整型一维数组 给定数组
 * @return int整型一维数组
*/
func solve( n int ,  m int ,  a []int ) []int {
    // 一维数组右移m个位置
	//数组拼接
	m = m%n
	arr1 := a[n-m:]
	a = a[:n-m]
	//将arr1放到前面
	arr1 = append(arr1, a...)
	return arr1
}