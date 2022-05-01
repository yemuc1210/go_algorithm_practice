package lt689
/*
689. 三个无重叠子数组的最大和
给你一个整数数组 nums 和一个整数 k ，
找出三个长度为 k 、互不重叠、且 3 * k 项的和最大的子数组，
并返回这三个子数组。

以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置（下标从 0 开始）。如果有多个结果，返回字典序最小的一个。
*/
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	sum := []int{}
	cur:=0
	for i:=0;i<k;i++{
		cur += nums[i]
	}
	sum = append(sum, cur)
	for i:=k;i<len(nums);i++{
		cur += nums[i] - nums[i-k]
		sum = append(sum, cur)
	}  //sum 长度k的子数组之和
	n := len(sum)
	left,right := make([]int,n),make([]int,n)
	for i:=range right{
		right[i] = n-1
	}
	for i:=1;i<n;i++{//left表示下标0-i的范围内sum[i]最大的下标i
		if sum[i] > sum[left[i-1]]{
			left[i] = i
		}else{
			left[i] = left[i-1]
		}
	}
	for i:=n-2;i>=0;i--{//right表示i-n-1范围内sum[i]最大
		if sum[i] >= sum[right[i+1]] {
			right[i] = i
		}else{
			right[i] = right[i+1]
		}
	}
	max := 0 
	res := make([]int,3)
	for i:=k;i<n-k;i++{//枚举中间坐标i
		if max < sum[i] + sum[left[i - k]] + sum[right[i + k]] {
			max =sum[i] + sum[left[i - k]] + sum[right[i + k]]
			res[0] = left[i-k]
			res[1] = i
			res[2] = right[i+k]
		}
	}
	return res
}