package NC106

import "sort"

/**
 * 最大乘积
 * @param A int整型一维数组
 * @return long长整型
 */
func solve( A []int ) int64 {
    // write code here
	sort.Ints(A)
	n := len(A)
	//两个负数+一个整数   三个整数
	a1 := int64(A[0]*A[1]*A[n-1])
	a2 := int64(A[n-1]*A[n-2]*A[n-3])
	if a1>a2 {
		return a1
	}
	return a2
}