package NC35

/**
 * min edit cost
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @param ic int整型 insert cost
 * @param dc int整型 delete cost
 * @param rc int整型 replace cost
 * @return int整型
*/

/*
描述
给定两个字符串str1和str2，再给定三个整数ic，dc和rc，分别代表插入、删除和替换一个字符的代价，
请输出将str1编辑成str2的最小代价。

数据范围：0≤∣str1∣,∣str2∣≤5000，0≤ic,dc,rc≤10000 
要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)

插入+删除的效果可以是替换的效果
插入
删除
dp[m][n] : str1[:m]  变为str2[:n]的代价
转移方程：
	//插入  删除   替换
	dp[i][j] = Min(dp[i][j-1]  ic, dp[i-1][j] dc, dp[i-1][j-1]  min{ic+dc, rc}   ) 
*/
func minEditCost( str1 string ,  str2 string ,  ic int ,  dc int ,  rc int ) int {
    // write code here
	m,n := len(str1),len(str2)
	dp := make([][]int,m+1)
	for i:=0;i<=m;i++{
		dp[i] = make([]int, n+1)
		dp[i][0] = i*dc  //// str1[i] 变成 str2[0], 删掉 str1[i], 需要 i 步操作
	}
	for j:=1;j<=n;j++{
		dp[0][j] = j*ic                     //str1[0] 变为str2[j]  需要J步 insert
	}

	rc = Min(rc, ic+dc)     //注释掉也没影响
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = Min(dp[i][j-1]+ic, Min(dp[i-1][j]+dc, dp[i-1][j-1]+rc)) 
			}
		}
	}
	return dp[m][n]
}

func Min(a,b int) int {
	if a<b {
		return a
	}
	return b
}