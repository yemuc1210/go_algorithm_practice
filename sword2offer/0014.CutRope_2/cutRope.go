package offer15

import "math"

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

大数取余，这是和上一题的区别！！！！！！

*/
//暴力法可以求，找出所有的可能组合，用map存 组合：乘积
/*
第二个是动态规划法，转移方程
	dp[n]:n能分割的最大
	dp[i]=max(dp[i], j*(i-j), j*dp[i-j])   两个数i,i-j     或j和更多的数dp[i-j]
*/
func cuttingRope(n int) int {
	dp := make([]int,n+1)
	dp[0],dp[1]=1,1
	for i:=1;i<=n;i++{
		for j:=1;j<i;j++{
			dp[i]=max(dp[i], j*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


/*
根据数学推导的结论，分为三段等长是最优的
切分规则：
最优：3，把绳子切分为多个长度为3的片段，留下的长度可能为0，1，2的情况
次优：2，最后一段绳子为2 保留不拆分
最差：1 最后一段为1，则把一份3+1 替换为2+2
算法流程：
当 n≤3 时，按照规则应不切分，但由于题目要求必须剪成 m>1 段，
因此必须剪出一段长度为 1的绳子，即返回 n - 1 。
当 n>3 时，求 n 除以 3的 整数部分 a 和 余数部分 b （即 n = 3a + b ），
并分为以下三种情况：
	当 b = 0时，直接返回 3^a
	当 b = 1时，要将一个 1 + 3转换为 2+2，因此返回 3^{a-1}* 4 
	当 b = 2 时，返回 3^a* 2
*/
func cuttingRope_ByMath(n int)int{
	if n<=3 && n>1{
		return n-1
	}
	a,b := n/3,n%3
	if b==0{
		return int(math.Pow(3,float64(a)))
	}
	if b==1{
		return int(math.Pow(3,float64(a-1))*4)
	}
	
	return int(math.Pow(3,float64(a))*2)
	
}


//本题的解  多了求余
func cuttingRope_greedy(n int)int{
	if n<4{
		return n-1
	}
	res:=1
	for n>4{
		res *=3
		res %= 1000000007;
		n-=3
	}
	return res*n% 1000000007
}
