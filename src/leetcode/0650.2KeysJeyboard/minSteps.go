package lt650

import "math"

/*

我们可以使用动态规划的思路解决本题。

设 f[i 表示打印出 i 个 A 的最少操作次数。由于我们只能使用「复制全部」和「粘贴」两种操作，
那么要想得到 i 个A，我们必须首先拥有 j 个 A，使用一次「复制全部」操作，
再使用若干次「粘贴」操作得到 i 个 A。
因此，这里的 j 必须是 i 的因数，「粘贴」操作的使用次数即为i/j-1
我们可以枚举 j 进行状态转移，这样就可以得到状态转移方程：
	f[i] = min{f[j]+i/j}

dp
	如果n是质数，则结果为n，只能一个个粘贴
	如果为合数，则结果是分解因式的结果之和
	dp[8]=dp[4]+dp[2]
*/

func minSteps(n int) int {
	dp := make([]int,n+1)
	for i:=2;i<=n;i++{
		dp[i] = i
		for j:=2;j<=int(math.Sqrt(float64(i)));j++{
			if i%j==0{
				dp[i]=dp[j]+dp[i/j]
				break
			}
		}
	}
	return dp[n]

}
