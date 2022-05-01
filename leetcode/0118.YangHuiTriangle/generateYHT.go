package lt0118

//杨辉三角   动态规划
func generate(numRows int)[][]int{
	if numRows<1 {
		return [][]int{}
	}
	ans := make([][]int, numRows)
    for i := range ans {
        ans[i] = make([]int, i+1)
        ans[i][0] = 1
        ans[i][i] = 1
        for j := 1; j < i; j++ {
            ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
        }
    }
    return ans
}