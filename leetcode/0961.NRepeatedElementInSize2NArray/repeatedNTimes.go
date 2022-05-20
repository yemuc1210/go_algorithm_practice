package main

import "fmt"

// len(nums) == 2*n
// 包含n+1个不同的元素
// 恰有一个元素重复n次
// 返回该元素
func repeatedNTimes(nums []int) int {
	// map
	// sort
	// 异或？不好用
	// 出现重复元素则就是
	// 1 0 2 0 n=2   一半的元素都是相等的
	// 隔开n个x，则需要n-1个额外的数，现有另外n个，所以多出一个，则
	// 最多使相邻x之间间隔两个额外的数字
	// 对于每一个nums[i] 判断前一个 前两个 前三个 判断是否重复
	n := len(nums)
	for i := 0; i < n; i++ {
		num := nums[i]
		if i-1 >= 0 && nums[i-1] == num {
			return num
		}
		if i-2 >= 0 && nums[i-2] == num {
			return num
		}
		if i-3 >= 0 && nums[i-3] == num {
			return num
		}
	}
	return -1
}

func main() {
	fmt.Println("vim-go")
	nums := []int{1, 2, 3, 3}
	res := repeatedNTimes(nums)
	fmt.Println(res)
}
