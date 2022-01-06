package main

import "fmt"

/*
描述
存在n+1个房间，每个房间依次为房间1 2 3...i，
每个房间都存在一个传送门，i房间的传送门可以把人传送到房间pi(1<=pi<=i),
现在路人甲从房间1开始出发(当前房间1即第一次访问)，每次移动他有两种移动策略：
    A. 如果访问过当前房间 i 偶数次，那么下一次移动到房间i+1；
    B. 如果访问过当前房间 i 奇数次，那么移动到房间pi；
现在路人甲想知道移动到房间n+1一共需要多少次移动；
输入描述：
第一行包括一个数字n(30%数据1<=n<=100，100%数据 1<=n<=1000)，表示房间的数量，
接下来一行存在n个数字 pi(1<=pi<=i), pi表示从房间i可以传送到房间pi。
输出描述：
输出一行数字，表示最终移动的次数，最终结果需要对1000000007 (10e9 + 7) 取模。
*/
/*
仔细分析 1<=pi<=i 知道用动态规划做。
记录第一次到达i为dp[i]，此时前面的所有门肯定是已经到达偶数次了
因为传送只会后退，前进的唯一方式是偶数次到达并+1，不能跳跃
所以到达i门前面所有门都走过并且经过偶数次（反正法也可以证明）
dp[i]=dp[i-1]+第二次到达i-1 + 1
第一次到达i-1门后再走一步会回到p[i-1]，此时p[i-1]门到达奇数次，其他所有门到达偶数次
这和第一次到达p[i-1]门的情况完全相同，所以从p[i-1]门回到i-1门，需要dp[i-1]-dp[p[i-1]]
所以dp[i] = dp[i-1] + dp[i-1] - dp[p[i-1]] + 1 + 1
dp[i] = 2 * dp[i-1] - dp[p[i-1]] + 2
*/
func directSolve(n int,p []int ) int {
	//需要记录访问次数，才能得到该使用哪种策略
	cntArr := make([]int,n+1)
	//根据描述，是先进入访问，此时cnt数值不改变，根据历史cnt决定策略
	var res int
	var curRoom int = 1  //当前第一个房间  序号[1,n]
	for curRoom != n+1{
		cntArr[curRoom] ++ 
		res =(res+1) % 1000000007 
		//根据cnt历史决定策略
		if cntArr[curRoom] % 2 == 0 {
			//偶数次  策略A	
			curRoom += 1
		}else if cntArr[curRoom]%2==1{
			curRoom = p[curRoom]
		}
		//有可能始终无法到达
	}
	return res%1000000007 
}
func getInput() (int, []int){
	var n int
	fmt.Scanln(&n)
	p := make([]int,n+1)
	for i:=1;i<n+1;i++{
		fmt.Scanf("%d%c",&p[i])
	}
	// fmt.Println(p)
	return n,p
}

func main() {
	n,p := getInput()
	// res := directSolve(n,p)
	res := dpSolve(n,p)
	fmt.Println( res)
}


//以上方法可能超时
func dpSolve(n int,p []int) int {
	dp := make([]int,n+2)
	var mod int = 1000000007
	for i := 2; i <= n+1; i++{
		dp[i] = (2 * dp[i-1] - dp[p[i-1]] + 2) % mod;
	}
	if dp[n]<0 {
		return dp[n+1]+mod
	}
	return dp[n+1]
}