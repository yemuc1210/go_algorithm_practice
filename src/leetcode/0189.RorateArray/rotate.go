package lt189 

/*
数组右移动
翻转递归
先翻转整体，再翻转[k,n-1]
*/
func rotate(nums []int, k int)  {
	k = k % len(nums) 
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func reverse(a []int){
	for i,n :=0,len(a); i<n/2 ;i++{
		a[i],a[n-i-1] = a[n-i-1],a[i]
	}
}