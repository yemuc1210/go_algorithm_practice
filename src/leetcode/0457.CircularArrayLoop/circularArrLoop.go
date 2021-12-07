package lt457

/*
nums[i] >0  前
nums[i] <0 后
因为是环形，所以第一个元素后移一个到最后一个；最后一个前移一个到第一个

判断是否存在循环。
使用map记录下标访问次数，超过一次就是存在？   不对

使用快慢指针判断循环，本题规定循环序列长度k>1，循环是同向的  nums[i]*nums[j]>0

若没有循环，则将访问过的路径置为0，表示访问过
*/
func circularArrayLoop(nums []int) bool {
	if len(nums) == 0{
		return false
	}
	for i:=0;i<len(nums);i++{
		if nums[i] == 0{
			continue //对应置零的操作
		}
		//快慢指针判断循环
		slow,fast,val := i, getNextIndex(nums,i),0
		for nums[fast]*nums[i] > 0 && nums[getNextIndex(nums,fast)]*nums[i]>0{   //同向  方向不能改变
			if slow == fast{
				//追上
				if slow == getNextIndex(nums,slow){
					break  //单元素loop
				}
				return true
			} 
			slow = getNextIndex(nums,slow)   //慢指针 一步
			fast = getNextIndex(nums,getNextIndex(nums,fast))   //快指针每次移动两步

		}
		slow,val = i,nums[i]
		for nums[slow]*val >0{
			next := getNextIndex(nums,slow)
			nums[slow]=0
			slow = next
		}
	}
	return false
}
func getNextIndex(nums []int, index int) int{
	//返回i
	// return (nums[i]+i)%len(nums)
	return ((nums[index]+index)%len(nums) + len(nums)) % len(nums)   //保证返回值在【0，len(nums)】
	
}