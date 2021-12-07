package lt131
/*lt131   回溯
剑指 Offer II 086. 分割回文子字符串
给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，
返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。
*/

/*
单个字符是回文
两个字符中  两个相同的是回文
三个

递归回溯  对字符串s 有len(s)种分割左右  加入左侧不是，舍弃，加入左侧是回文，则对右侧递归
*/
func partition(s string) (ans [][]string) {
    n := len(s)
    /*
	dp状态转移方程判断但回文
	f(i,j) = True  i>=j  空串  i==j长度为1
		= f(i+1,j-1) & (s[i]==s[j])   首尾相同且s[i+1,j-1]为回文串
	*/
	f := make([][]bool, n)

    for i := range f {
        f[i] = make([]bool, n)
        for j := range f[i] {
            f[i][j] = true   //f(i,j)标识s（i,j)是回文
        }
    }
    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            f[i][j] = s[i] == s[j] && f[i+1][j-1]
        }
    }
	/*
	假设当前搜索到字符串的第i个字符，且s[0...i-1]位置所有字符已经被分割成若干回文串
	此时需要枚举下一个回文串的右边界j，使得s[i..j]是一个回文串
	i开始，枚举j，判断s[i...j]是否是回文串，若是则加入结果数组，并以j+1作为新的i进行下一层搜索

	*/
    splits := []string{}
    var dfs func(int)
    dfs = func(i int) {
        if i == n {  //搜到最后字符串
            ans = append(ans, append([]string(nil), splits...))
            return
        }
        for j := i; j < n; j++ {
            if f[i][j] {//如果子串s[i,,,j]是回文串
                splits = append(splits, s[i:j+1])  //添加一个回文串
                dfs(j + 1) //下一层搜索
                splits = splits[:len(splits)-1]  //返回上一层，回溯
            }
        }
    }
    dfs(0)
    return
}


