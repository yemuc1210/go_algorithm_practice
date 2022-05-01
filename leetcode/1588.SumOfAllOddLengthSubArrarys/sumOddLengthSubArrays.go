package lt1588
/*
每日一题
给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。

子数组 定义为原数组中的一个连续子序列。

请你返回 arr 中 所有奇数长度子数组的和 


从当前元素出发，数组长度每次+2，之前的和累加
*/
func sumOddLengthSubarrays(arr []int) int {
	sum,tmp := 0,0
	for i:=0;i<len(arr);i++{
		sum += arr[i]
		tmp = arr[i]
		for j:=i;j+2<len(arr); j=j+2 {
			tmp = tmp + arr[j+1] + arr[j+2]
			sum += tmp
		}
	}
	return sum
}