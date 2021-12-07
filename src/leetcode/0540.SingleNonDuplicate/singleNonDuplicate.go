package lt540
//offerS70
/*
有序数组中，只有一个数字只出现一次。。。 lt 136  260
*/

/*
剑指 Offer II 070. 排序数组中只出现一次的数字
给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。

进阶: 采用的方案可以在 O(log n) 时间复杂度和 O(1) 空间复杂度中运行吗？
则不能遍历异或

二分查找，而数组只会是奇数的，因为只有一个出现一次，其他的都出现两次  len(num)>=3

移除中间一对，左右也是奇数
*/
func singleNonDuplicate(nums []int) int {
	low,high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		//判断中间元素是哪个  左侧子数组中的最后一个元素或右侧子数组中的第一个元素
		var halvesAreEven bool = false
		if (high-mid)%2 == 0{   //右半部分是奇数
			halvesAreEven = true
		}
		if nums[mid+1] == nums[mid] {
			if halvesAreEven {  //在奇数部分找
				low = mid + 2
			}else{
				high = mid - 1
			}
		}else if nums[mid-1] == nums[mid]{
			if halvesAreEven {
				high = mid - 2
			}else{
				low = mid+1
			}
		}else{
			return nums[mid]
		}
	}
	return nums[low]
}
