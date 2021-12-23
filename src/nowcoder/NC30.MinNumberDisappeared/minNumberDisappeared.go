package NC30

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 缺失的第一个正整数   数组 哈希  lt41 firstMissingPositive
 * 给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数
 * @param nums int整型一维数组 
 * @return int整型
*/
func minNumberDisappeared( nums []int ) int {
    // write code here
	n := len(nums)
	//超出范围
	for i:=range nums{
		if nums[i]<0 || nums[i] > n{
			nums[i]=0
		}
	}
	//如果某个数出现，比如出现了2，那么就把num[1]加上（len+1），这样相当于打了标记，上一步置为0也是为了避免原数组中本来就有>len的数干扰
    for i:=0;i<n;i++{
		tmp := nums[i] % (n+1)
		if tmp==0{
			continue
		}
		if nums[tmp-1] <= n {
			nums[tmp-1] += n+1
		}
	}  
	//遍历数组 找到第一个num[i]不大于len，则i+1未出现在原数组
	for i:=0;i<n;i++{
		if nums[i] < n+1{
			return i+1
		}
	}
	return n+1
}