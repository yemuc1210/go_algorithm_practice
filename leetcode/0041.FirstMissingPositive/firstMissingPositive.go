package lt41
/*困难
41. 缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：
输入：nums = [1,2,0]
输出：3

示例 2：
输入：nums = [3,4,-1,1]
输出：2

示例 3：
输入：nums = [7,8,9,11,12]
输出：1

*/
/*
二进制 是否有解决方案  相邻整数之间的二进制数位
0000
0001
0010
0011

求和公式 找到最大值 最小值
	最大值-最小值+1 = 个数   
	如果个数 ！= len(nums)  false

*/
func firstMissingPositive(nums []int) int {
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


//评论区copy
// def firstMissingPositive(self, nums: List[int]) -> int:
//         for j in range(1, 2 ** 31 - 1):
//             if j not in nums:
//                 return j