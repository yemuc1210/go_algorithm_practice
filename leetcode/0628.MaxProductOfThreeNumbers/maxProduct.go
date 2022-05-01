package lt628
import "sort"
//三个数的最大乘积
func maximumProduct(A []int) int64 {
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
