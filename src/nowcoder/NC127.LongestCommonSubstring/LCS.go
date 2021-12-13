package NC127

/**
 * longest common substring
 * 最长公共子串问题 需要构建子串
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
*/
func LCS( s1 string ,  s2 string ) string {
    lenLCCS,indexOfLCCSs1,indexOfLCCSs2 := 0,0,0
    m,n := len(s1),len(s2)
    dp := make([][]int,m+1)
    for i := 0 ; i <= m ; i++ {
        dp[i] = make([]int,n+1)
    }
    for i := 1 ; i <= m ; i++ {
        for j := 1 ; j <= n ; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
                if dp[i][j] > lenLCCS {
                    lenLCCS = dp[i][j]
                    indexOfLCCSs1 = i
                    indexOfLCCSs2 = j
                }
            }
        }
    }
    var res string
    for i,j := indexOfLCCSs1,indexOfLCCSs2 ; dp[i][j] > 0 ; i,j = i-1,j-1{
        res = string(s1[i-1]) + res
    }
    return res
}