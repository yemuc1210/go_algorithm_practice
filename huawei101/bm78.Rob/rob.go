package main
import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param nums int整型一维数组 
 * @return int整型
*/
func rob( nums []int ) int {
    // write code here
    if len(nums) == 0 {
        return 0
    }
    n := len(nums)
    if n==1 {
        return nums[0]
    }
    // dp[i] 抢劫[0,....,i]的最大价值
    dp := make([]int, n+1)
    dp[0] = nums[0]
    dp[1] = max(nums[0],nums[1])
    for i:=2;i<n;i++{
        dp[i] = max(dp[i-1], nums[i] + dp[i-2])
    }
    return dp[n-1]
}
func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}

func main() {
	nums := []int{1,2,3,4}
	res := rob(nums)
	fmt.Println(res)
}