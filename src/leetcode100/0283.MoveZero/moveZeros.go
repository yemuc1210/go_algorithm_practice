package lt283
/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

移动非零，补零
*/
func moveZeroes(nums []int)  {
	//遇到0跳过
	index := 0
	for i:=0;i<len(nums);i++{
		if nums[i] != 0 {
			nums[index] = nums[i]
			index+=1
		}
	}
	//将剩余补零
	for index < len(nums) {
		nums[index] = 0
		index ++
	}
}