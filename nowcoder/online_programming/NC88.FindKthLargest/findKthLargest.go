package NC88
/**
 * 按照快排的思路 找出第k大
 * @param a int整型一维数组 
 * @param n int整型 
 * @param K int整型 
 * @return int整型
*/
func findKth( a []int ,  n int ,  K int ) int {
    // write code here
	if len(a) == 0 {
		return 0
	}
	return selection(a, 0, n-1, len(a)-K)  //第k大
}

/*
快排每次确定一个元素的最终位置 根据partition结果 
可以知道kth元素位置
*/
func partition(a []int, low,high int) int {
	pivot := a[high]
	i := low - 1
	//开始划分
	for j:=low;j<high;j++{
		if a[j] < pivot {  
			i++
			a[j],a[i] = a[i],a[j]
		}
	}
	a[i+1],a[high] = a[high],a[i+1]  ////a[i]左边都是小于它的
	return i+1
}

func selection(arr []int, l, r, k int) int {
	if l == r {
		return arr[l]
	}
	p := partition(arr, l, r)
	if k == p {
		return arr[p]
	} else if k < p {
		return selection(arr, l, p-1, k)
	} else {
		return selection(arr, p+1, r, k)
	}
}