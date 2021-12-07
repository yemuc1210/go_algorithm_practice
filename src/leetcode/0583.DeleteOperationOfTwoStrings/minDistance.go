package lt583

/*
算法计划　　ｄｐ    1143  583

给定两个单词 word1 和 word2，
找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

给定两个字符串word1 和word2，分别删除若干字符之后使得两个字符串相同，则剩下的字符为两个字符串的公共子序列。
为了使删除操作的次数最少，剩下的字符应尽可能多。当剩下的字符为两个字符串的最长公共子序列时，删除操作的次数最少。
因此，可以计算两个字符串的最长公共子序列的长度，然后分别计算两个字符串的长度和最长公共子序列的长度之差，
即为两个字符串分别需要删除的字符数，两个字符串各自需要删除的字符数之和即为最少的删除操作的总次数。


*/

func minDistance(word1 string, word2 string) int {
    l1:=len(word1)
    l2:=len(word2)
    l:=longestCommonSubsequence(word1,word2)
    return l1+l2-l*2
}

func longestCommonSubsequence(text1 string, text2 string) int {
    m:=len(text1)
    n:=len(text2)
    dp:=make([][]int,m+1)
    for i:=range dp{
        dp[i]=make([]int,n+1)
    }
    for i:=1;i<=m;i++{
        for j:=1;j<=n;j++{
            if text1[i-1]==text2[j-1]{
                dp[i][j]=dp[i-1][j-1]+1
            }else{
                dp[i][j]=max(dp[i-1][j],dp[i][j-1])
            }
        }
    }
    return dp[m][n]
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
