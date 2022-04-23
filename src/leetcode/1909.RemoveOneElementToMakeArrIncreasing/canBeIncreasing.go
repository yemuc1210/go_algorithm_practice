package main

import "fmt"

// 如果恰好删除一个元素后，数组严格递增，则true
func canBeIncreasing(nums []int) bool {
	// 检查删去下标为idx之后的元素是否严格递增
	check := func(idx int) bool {
		var preMax = -1
		var count int
		n := len(nums)
		for i := 0; i < n; i++ {
			if i == idx  {   // 跳过下标idx
				continue
			}
			// 因为是递增，每次仅需记住上一个最大值
			// 当当前>上一个最大，计数
			if preMax == -1 || preMax < nums[i] {
				preMax = nums[i]
				count++
			}
		}
		// 判断递增数量是否刚好等于n-1
		return count == n-1
	}
	

	for i:=1;i<len(nums);i++{
		if nums[i]<=nums[i-1]{   // 
			return check(i-1) || check(i)	//只有都是false才会返回false
		}
	}
	return true
}


func main() {
	nums := []int{1,2,10,5,7}
	res := canBeIncreasing(nums)
	fmt.Println(res)
}