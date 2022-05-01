package main

import (
    "fmt"
)

func main(){
    var n int
    fmt.Scanln(&n)
    dp := make([][]int,n)
    for i:=range dp{
        dp[i]=make([]int,n)
    }
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            fmt.Scanf("%d%c",&dp[i][j])
        }
    }
    //机器人问题
    //第一行
    for i:=1;i<n;i++{
        dp[i][0] += dp[i-1][0] 
    }
    //第一列
    for j:=1;j<n;j++{
        dp[0][j] += dp[0][j-1]
    }
    for i:=1;i<n;i++{
        for j:=1;j<n;j++{
            dp[i][j] += min(dp[i-1][j], dp[i][j-1])
        }
    }
    fmt.Println(dp[n-1][n-1])
}

func min(a,b int)int {
    if a<b {
        return a
    }
    return b
}