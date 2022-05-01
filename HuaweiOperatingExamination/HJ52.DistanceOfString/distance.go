package main

import (
    "fmt"
)

func main() {
    //本题有多组输入数据
    var s1,s2 string
    var n int
    var err error
    for {
        n,err = fmt.Scanln(&s1)
        if n==0 {
            break
        }
        if err!= nil {
            panic(err)
        }
        _,err := fmt.Scanln(&s2)
        if err!=nil {
            panic(err)
        }
        //下面求解
        minDistance := solve(s1,s2)
        fmt.Println(minDistance)
    }
}

// 允许的操作：插入 替换 删除
//由 src 转换成 dst 的最小操作次数
func solve(src,dst string) int {
    //dp[i][j] 表示 src[:i] 变成dst[:j]的
    m,n := len(src),len(dst)
    dp := make([][]int, m+1)
    for i:=range dp {
        dp[i]= make([]int,n+1)
    }
    for i:=0;i<m+1;i++{
        dp[i][0] = i  //src[i] 变成 dst[0] ，需要进行删除操作
    }
    for j:=0;j<n+1;j++{
        dp[0][j] = j   //插入操作
    }
    
    for i:=1;i<m+1;i++{
        for j:=1;j<n+1;j++{
            // 状态转换公式：
            if src[i-1] == dst[j-1] {
                //src[i-1] dst[j-1]都是相同的 之前的操作将两个字符串之前的部分变成相同的
                dp[i][j] = dp[i-1][j-1]
            }else{
                // Min{插入 删除 替换}
                dp[i][j] = Min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}

func Min(args ...int) int {
    min := args[0]
    for _,item := range args {
        if item < min {
            min = item
        }
    }
    return min
}

/*
分析：
有些操作是等价的，比如A插入和B删除

dp[i][j] word1前i个字符 变换到word2前j个字符 的最少操作次数
增 dp[i][j] = dp[i][j-1] + 1   
删  dp[i][j] = dp[i-1][j] + 1
改 dp[i][j] = dp[i-1][j-1] + 1
*/ 