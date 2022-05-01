package NC48

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 在旋转过的有序数组中寻找目标值
 * 简单题 lt33
 * @param nums int整型一维数组 
 * @param target int整型 
 * @return int整型
*/
func search(nums []int, target int) int {
	length := len(nums)
	left,right := 0, length-1

	for left <= right {
		mid := (left + right)/2
		//判断中间数和最右
		if nums[mid] == target {
			return mid
		}else if nums[mid] < nums[right] {
			//右半段有序
			if nums[mid] < target && target <= nums[right] {
				//说明target在右半段
				left = mid+1
			}else{
				right = mid-1
			}
		}else{
			//右侧不一定有序
			if nums[left]<=target && target < nums[mid]{
				right = mid-1
			}else{
				left = mid+1
			}
		}
	}
	return -1
}