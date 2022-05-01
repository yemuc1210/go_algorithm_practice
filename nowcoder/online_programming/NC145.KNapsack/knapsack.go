package NC145

//01背包

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算01背包问题的结果
 * @param V int整型 背包的体积
 * @param n int整型 物品的个数
 * @param vw int整型二维数组 第一维度为n,第二维度为2的二维数组,vw[i][0],vw[i][1]分别描述i+1个物品的vi,wi
 * @return int整型
*/
func knapsack( V int ,  n int ,  vw [][]int ) int {
    // write code here
	dp := make([][]int,n)
	for i:=range dp{
		dp[i] = make([]int, V)
	}
	for i := 0; i<n;i++ {  //
 		for j := 0;j<V;j++{
			if j<vw[i-1][0] {  //体积-种类
				//dp[i][j]表示：对于前i个物品，当前背包的容量为j时，这种情况下可以装下的最大价值是dp[i][j]。
				dp[i][j] = dp[i-1][j];     //背包容量j不够存 
			}else{
				dp[i][j] = Max(dp[i-1][j],  dp[i-1][j-vw[i-1][0]]+vw[i-1][1]);
			}
		}
	}
	return dp[n-1][V-1]
}
/*
 for(int i = 1; i<=n;i++){
     for(int j = V;j>0;j--){
         if(j<vw[i-1][0]){
             dp[j] = dp[j];
         }else{
             dp[j] = Math.max(dp[j],dp[j-vw[i-1][0]]+vw[i-1][1]);
         }
     }
 }
*/

func Max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

/*
描述
已知一个背包最多能容纳体积之和为v的物品
现有 n 个物品，第 i 个物品的体积为 vi , 重量为 wi
求当前背包最多能装多大重量的物品?
数据范围： 1≤v≤1000 ,1≤n≤1000 ，1≤vi≤1000 ，1≤wi≤1000

进阶 ：O(n⋅v)
*/