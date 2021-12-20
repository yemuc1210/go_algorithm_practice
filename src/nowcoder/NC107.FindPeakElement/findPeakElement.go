package NC107

// import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 寻找峰值
 * 数组包含多个峰值，返回任意一个位置即可
 * @param nums int整型一维数组 
 * @return int整型
*/
func findPeakElement( nums []int ) int {
    // 峰顶严格大于两侧
// 规律
// 规律一：如果nums[i] > nums[i+1]，则在i之前一定存在峰值元素
// 规律二：如果nums[i] < nums[i+1]，则在i+1之后一定存在峰值元素
	left,right :=0,len(nums)-1
	for left<right{
		mid := (left+right)/2
		if nums[mid] < nums[right]{
			left = mid+1
		}else{
			right = mid-1
		}
	}
	return left
}