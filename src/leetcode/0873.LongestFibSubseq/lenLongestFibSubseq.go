package lt873

/*  lt873
剑指 Offer II 093. 最长斐波那契数列
如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：

n >= 3
对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。

（回想一下，子序列是从原序列  arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）
*/

/*
dp
根据斐波那契数列定义，需要两位而不是一位才能定义一个数列。
因此，设dp[i][j]表示以A[i],A[j]结尾且i < j 且 A[i] <A[j]的最长斐波那契式的子序列的长度。
初始值：
	任意两个元素A[i],A[j]之间最小为2，dp[i][j] = 2
转移方程：
	当arr[k]+arr[i]=arr[j]那么dp[i][j]=max(dp[i][j],dp[k][i]+1)

*/
func lenLongestFibSubseq(arr []int) int {
    var res int
    var m=len(arr)
    if m<3 {
        return res
    }
    var dp=make([][]int,m-1)//dp[i][j], because i<j and j<m so i<m-1
    for i:=range dp{
        //初始化
        for j:=0;j<m;j++{
            dp[i]=append(dp[i],2)
        }
    }
    var valueMapIndex=make(map[int]int)//arr值及对应的下标
    for j:=0;j<m;j++{
        //数组严格递增
        valueMapIndex[arr[j]]=j
        for i:=0;i<j;i++{
            t:=arr[j]-arr[i]
            if k,ok:=valueMapIndex[t];ok==true&&k<i{
                //说明存在arr[k]+arr[i]=arr[j]且k<i
                dp[i][j]=max(dp[i][j],dp[k][i]+1)
            }
            res=max(res,dp[i][j])
        }
    }
    if res==2{
        //数列肯定大于2
        res=0
    }
    return res
}

func max(x,y int)int{
    if x<y {
        return y
    }
    return x
}

