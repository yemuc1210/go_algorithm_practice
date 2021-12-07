package lt926

/*
剑指 Offer II 092. 翻转字符
如果一个由 '0' 和 '1' 组成的字符串，
是以一些 '0'（可能没有 '0'）后面跟着一些 '1'（也可能没有 '1'）的形式组成的，
那么该字符串是 单调递增 的。

我们给出一个由字符 '0' 和 '1' 组成的字符串 s，
我们可以将任何 '0' 翻转为 '1' 或者将 '1' 翻转为 '0'。

返回使 s 单调递增 的最小翻转次数。
*/
func minFlipsMonoIncr(s string) int {
	n := len(s)
	var tail0, tail1 int
	if s[0] == '0' {
		tail1 = 1
	} else {
		tail0 = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			tail1 = min(tail0, tail1) + 1
		} else {
			tail1 = min(tail0, tail1)
			tail0++
		}
	}
	return min(tail0, tail1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


