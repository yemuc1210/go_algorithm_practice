package leetcode
/*
小朋友 A 在和 ta 的小伙伴们玩传信息游戏，游戏规则如下：

有 n 名玩家，所有玩家编号分别为 0 ～ n-1，其中小朋友 A 的编号为 0
每个玩家都有固定的若干个可传信息的其他玩家（也可能没有）。
传信息的关系是单向的（比如 A 可以向 B 传信息，但 B 不能向 A 传信息）。
每轮信息必须需要传递给另一个人，且信息可重复经过同一个人
给定总玩家数 n，以及按 [玩家编号,对应可传递玩家编号] 关系组成的二维数组 relation。
返回信息从小 A (编号 0 ) 经过 k 轮传递到编号为 n-1 的小伙伴处的方案数；
若不能到达，返回 0。
*/
func numWays(n int, relation [][]int, k int) int {
	dp := make([][]int,k+1)
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int, n)
	}
	dp[0][0]=1   //第0轮到0，总方案为1
	for i:=1;i<=k;i++{
		for  _,r := range relation{
			dp[i][r[1]] += dp[i-1][r[0]]
		}
	}
	return dp[k][n-1]
}

/*
dfs 也能做
从起始为0的开始，dfs中记录迭代步数，并且如果是遇到n-1，则结果加1
*/
func numWays_DFS(n int,relation [][]int,k int)int{
	res := 0	
	for i:=0;i<len(relation);i++{
		if relation[i][0]==0{//从0开始做dfs
			dfs(k,n,relation,1,relation[i][1],&res)
		}
	}
	return res
}

func dfs(k,n int, relation [][]int,step,start int,ans *int){
	//递归终止条件
	if step==k{//迭代k步
		if start==n-1{
			*ans = *ans + 1
		}
		return 
	}
	for i:=0;i<len(relation);i++{
		if relation[i][0]==start{
			// dfs(k,n,relation,step+1,relation[i][1],&ans)//  &ans错误
			dfs(k,n,relation,step+1,relation[i][1],ans)
		}
	}
}