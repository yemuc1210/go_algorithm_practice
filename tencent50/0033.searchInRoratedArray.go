package tencent50


/* z中等
33. 搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，
   如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
*/
/*
一维旋转  移位的感觉
找到旋转的节点 旋转节点左右两个数是逆序的

二分吧
*/
func search(nums []int, target int) int {
	povit := 0 
	//遍历一遍找到pivot
	for povit < len(nums)-1 {
		if nums[povit] > nums[povit+1] {
			//逆序
			break
		}
		povit++
	}
	//在【0：pivot] [pivot+1:] 两边 二分查找
	leftRes := binarySearch(nums,0,povit,target)
	if leftRes != -1{
		return leftRes
	}else{
		return binarySearch(nums,povit+1, len(nums)-1, target)
	}
}

func binarySearch(nums []int, start,end int, target int) int{
	for start <= end {
		mid := start +(end - start)/2
		if nums[mid] == target{
			return mid
		}else if nums[mid] > target {
			end = mid-1
		}else{
			start = mid+1
		}
	}
	return -1
}