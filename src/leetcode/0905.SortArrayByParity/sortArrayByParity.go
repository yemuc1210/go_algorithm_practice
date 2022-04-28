package main

import "fmt"

// 奇偶排序数组
// 双指针
func sortArrayByParity(nums []int) []int {
	front, rear := 0, len(nums)-1 // front-数组前面，偶数  rear-尾部 奇数
	for front < rear {            // front==rear时 应该跳出
		// front指向第一个奇数
		for front < rear && nums[front]%2 == 0 {
			front++
		}
		// rear 指向偶数
		for front < rear && nums[rear]%2 == 1 {
			rear--
		}
		// 判断是否可以交换
		if front >= rear {
			break
		}
		// 交换
		nums[front], nums[rear] = nums[rear], nums[front]
	}
	return nums
}
func main() {
	fmt.Println("vim-go")
	nums := []int{3, 1, 2, 4}
	res := sortArrayByParity(nums)
	fmt.Println(res)
}
