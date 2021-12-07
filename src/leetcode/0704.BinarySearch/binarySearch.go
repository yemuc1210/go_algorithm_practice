package lt704

func search(nums []int, target int)int{
	//题目假设nums中所有元素是不重复的   升序
	left, right := 0,len(nums)-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] > target{
			//左侧
			right = mid-1
		}else if nums[mid] < target{
			//右侧
			left = mid + 1
		}else{
			return mid
		}

	}
	return -1
}