package offerS101
/*lt416  dp
剑指 Offer II 101. 分割等和子集
给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。

背包
*/
// d(i, s) : 是否存在：nums区间[0, i] 中取一些元素，使其和为s
// d(i, s) = d(i-1, s){不取nums[i]} || d(i-1, s-nums[i]){取nums[i]}
// max(i) = nums.size()-1
// max(s) = sum(nums)/2

func canPartition(nums []int) bool {
	sum := 0
	for _,v := range nums{
		sum +=v
	}
	if sum & 1 == 1{  //奇数
		return false
	}
	sum = sum>>1
	m,n := len(nums),sum   //sum=sum/2  分为两部分嘛
	dp := make([][]bool, m+1)
	for i:=range dp{
		dp[i] = make([]bool, n+1)
		dp[i][0] = true    //在区间[0,i]中取某些元素 和为0   true:不取即可
	}
	for i := 1; i <= len(nums); i++ {
        for j := 1; j <= sum; j++ {
            if j - nums[i-1] >= 0 {  //
                dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
	return dp[m][sum]

}