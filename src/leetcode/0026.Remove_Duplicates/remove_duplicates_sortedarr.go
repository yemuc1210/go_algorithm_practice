/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/
package leetcode

// import "fmt"

//原地删除，通过移位实现
func removeDuplicates(nums []int) int {
	if len(nums) == 1{
		return 1
	}
	idx := 0
	lst_idx := len(nums)
	for i:=0;i<lst_idx;i++{
		cnt := 0   //移位长度
		for idx=i+1;idx<lst_idx;idx++{
			//找到第一个不等于nums[i]的元素
			if nums[idx] != nums[i]{
				break
			}else{
				cnt ++
			}
		}
		// fmt.Println(idx)
		//移位
		for k:=idx;k<lst_idx;k++{
			nums[k-cnt] = nums[k]
		}
		lst_idx -= cnt
	}
	return lst_idx
}