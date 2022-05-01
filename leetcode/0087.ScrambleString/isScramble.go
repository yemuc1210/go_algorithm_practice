package lt87
/*困难
扰乱字符串
87. 扰乱字符串
使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
	如果字符串的长度为 1 ，算法停止
	如果字符串的长度 > 1 ，执行下述步骤：
		在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
		随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y + x 。
		在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。
*/
func isScramble(s1, s2 string) bool {
    n := len(s1)
    dp := make([][][]int8, n)
    for i := range dp {
        dp[i] = make([][]int8, n)
        for j := range dp[i] {
            dp[i][j] = make([]int8, n+1)
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }

    // 第一个字符串从 i1 开始，第二个字符串从 i2 开始，子串的长度为 length
    // 和谐返回 1，不和谐返回 0
    var dfs func(i1, i2, length int) int8
    dfs = func(i1, i2, length int) (res int8) {
        d := &dp[i1][i2][length]
        if *d != -1 {
            return *d
        }
        defer func() { *d = res }()

        // 判断两个子串是否相等
        x, y := s1[i1:i1+length], s2[i2:i2+length]
        if x == y {
            return 1
        }

        // 判断是否存在字符 c 在两个子串中出现的次数不同
        freq := [26]int{}
        for i, ch := range x {
            freq[ch-'a']++
            freq[y[i]-'a']--
        }
        for _, f := range freq[:] {
            if f != 0 {
                return 0
            }
        }

        // 枚举分割位置
        for i := 1; i < length; i++ {
            // 不交换的情况
            if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
                return 1
            }
            // 交换的情况
            if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
                return 1
            }
        }

        return 0
    }
    return dfs(0, 0, n) == 1
}
