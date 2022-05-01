package offer03

/*

剑指 Offer 03. 数组中重复的数字
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，
但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。


异或一遍  自身与自身异或为0   0与任何数异或为其本身   最后结果是不重复的。。。

关键点：长度为n，里面数字范围是0~n-1，所以可以使用原地交换，使得Nums[i]=i
*/

func findRepeatNumber(nums []int) int {
	//特殊判断
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	i:=0
	for i<len(nums){
		if nums[i]==i{
			i++
			continue  //无需交换
			//方法二的解法有一个关键点是只有 nums[i] == i 的时候i才递增，这样保证找到相同元素前不会漏掉某些元素的处理
		}else if nums[i] == nums[nums[i]] {
			//索引Nums[i] 与 索引i 元素相同，找到重复值
			return nums[i]
		}else{//交换索引为 ii 和 nums[i]nums[i] 的元素值，将此数字交换至对应索引位置。
			nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
		}
	}
	return -1
}