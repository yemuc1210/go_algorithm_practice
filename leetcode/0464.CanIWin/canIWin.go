package main

// 原100ganem: 轮流选择1-10的任意整数，累计整数和 先达到或超过100的赢
// 现在：1-15 不放回
// 给定可选择的最大数，累计和
// 若先出手的稳赢，则返回true。假设两位玩家都表现最佳
// func canIWin(n int, t int) bool {
// 	choosenState := make([]bool, n+1)  // 数字被选择的情况
// 	// dp[statue][k]: 第k轮，选择数字state，原始回合的先手能否获胜  k从0开始，根奇偶判断先后手
// 	dp := make([][]int, 1<<20)
// 	var dfs func(state,tot,k int) int
// 	dfs = func(state, total, k int) int {
// 		// 表示当前为state,累计和total,轮数k
// 		// 转移过程中，如果当前回合的决策，可以使累加和超过t，则获胜
// 		// 或者，当前决策，导致下一回合的玩家失败，当前回合获胜
// 		// 否则失败
// 		if state == (1<<n-1) && total < t {
// 			return -1
// 		}
// 		if dp[state][k%2]!=0 {
// 			return dp[state][k%2]
// 		}//记忆化搜索
// 		var hope int
// 		if k%2 ==0 {hope = 1}else{hope=-1}
// 		for i:=0;i<n;i++{
// 			if state>>i & 1 == 1 {
// 				continue
// 			}
// 			if total+i+1 >=t {
// 				dp[state][k%2]=hope
// 				return dp[state][k%2]
// 			}
// 			if dfs(state|(1<<i), total+i+1,k+1) == hope {
// 				dp[state][k%2] = hope
// 				return dp[state][k%2]
// 			} 
// 		}
// 		dp[state][k%2] = -hope
// 		return dp[state][k%2]
// 	}
// 	if t==0 {return true}
// 	return dfs(0,0,0) == 1
// }
func canIWin(maxChoosableInteger, desiredTotal int) bool {
    if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
        return false
    }

    dp := make([]int8, 1<<maxChoosableInteger)
    for i := range dp {
        dp[i] = -1
    }
    var dfs func(int, int) int8
    dfs = func(usedNum, curTot int) (res int8) {
        dv := &dp[usedNum]
        if *dv != -1 {
            return *dv
        }
        defer func() { *dv = res }()
        for i := 0; i < maxChoosableInteger; i++ {
            if usedNum>>i&1 == 0 && (curTot+i+1 >= desiredTotal || dfs(usedNum|1<<i, curTot+i+1) == 0) {
                return 1
            }
        }
        return
    }
    return dfs(0, 0) == 1
}
