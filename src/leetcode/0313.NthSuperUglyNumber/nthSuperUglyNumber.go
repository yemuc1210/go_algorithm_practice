package lt313

import "math"

/*
超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。

质因数只是出现在primes数组中，并不规定存在几个
题目数据 保证 primes[i] 是一个质数
primes 中的所有值都 互不相同 ，且按 递增顺序 排列

质数由只能是2、3、5扩展到k个数组元素

动态规划
264题变种
dp[i]表示第i个超级丑数
dp[1] = 1
参考264做法，创建一个指针数组pointers，则下一个超级丑数=当前指针的超级丑数乘以对应的质因数

2《=i《=n  dp[i]= min{dp[ pointers[j] * primes[j]]} 对每个0<=j<=m，分别比较dp[i]与dp[pointers[j]]是否相等，若相等则pointers[j]+1  （倍数增加）


*/
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m := len(primes)
	pointers := make([]int, m)
	for i:= range pointers {
		pointers[i] = 1
	}

	for i:=2; i<=n; i++ {
		nums := make([]int, m)
		minNum := math.MaxInt64
		for j, p := range pointers {
			nums[j] = dp[p] * primes[j]
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j,num := range nums{
			if minNum == num{
				pointers[j]++
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
