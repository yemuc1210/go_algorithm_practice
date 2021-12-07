package NC22

/**
 * 合并两个有序的数组  升序
 * @param A int整型一维数组 
 * @param B int整型一维数组 
 * @return void
*/
func merge( A []int ,  m int, B []int, n int )  {
    // write code here
	//假设A空间足够大  思路：从后往前
	for i,j,k:= m-1,n-1,m+n-1;i>=0||j>=0;k-- {
		//可能出现某个指针先到0
		var cur int
		if i == -1{
			cur = B[j]
			j--
		}else if j==-1{
			cur = A[i]
			i--
		}else if A[i] > B[j]{
			cur = A[i]
			i--
		}else{
			cur = B[j]
			j--
		}
		A[k] = cur
	}
}