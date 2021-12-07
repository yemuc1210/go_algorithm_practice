package NC68

/**
 * 跳台阶问题 青蛙一次跳一个或两个
 * @param number int整型 
 * @return int整型
*/
func jumpFloor( number int ) int {
    // write code here
	dp := make([]int, number+1)   //dp[i]跳number个有几种跳法
	if number <= 2{
		return number
	}
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i:=3;i<=number;i++{
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}