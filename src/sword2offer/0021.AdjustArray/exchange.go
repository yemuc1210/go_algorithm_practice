package offer21

/*

剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，
所有偶数位于数组的后半部分。

 

示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4] 
注：[3,1,2,4] 也是正确的答案之一。
*/

func exchange(nums []int) []int {
	//双指针
	front,rear := 0,len(nums)-1    //前奇数后偶数
	for front < rear {
		//寻找到第一个偶数
		for front<rear && nums[front]%2!=0{
			front ++
		}
		for front<rear && nums[rear]%2==0{
			rear--
		}
		//交换
		nums[front],nums[rear] = nums[rear],nums[front]
	}
	return nums
}