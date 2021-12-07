package lt680
/*
680. 验证回文字符串 Ⅱ
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

在允许最多删除一个字符的情况下，同样可以使用双指针，通过贪心实现。
初始化两个指针 low 和 high 分别指向字符串的第一个字符和最后一个字符。
每次判断两个指针指向的字符是否相同，
如果相同，则更新指针，将 low 加 1，high 减 1，

然后判断更新后的指针范围内的子串是否是回文字符串。
如果两个指针指向的字符不同，则两个字符中必须有一个被删除，
此时我们就分成两种情况：
	即删除左指针对应的字符，留下子串s[low+1:high]，
	或者删除右指针对应的字符，留下子串 s[low:high−1]
当这两个子串中至少有一个是回文串时，就说明原始字符串删除一个字符之后就以成为回文串。

*/

func validPalindrome(s string) bool {
    low, high := 0, len(s) - 1
    for low < high {
        if s[low] == s[high] {
            low++
            high--
        } else {
            flag1, flag2 := true, true
            for i, j := low, high - 1; i < j; i, j = i + 1, j - 1 {
                if s[i] != s[j] {
                    flag1 = false
                    break
                }
            }
            for i, j := low + 1, high; i < j; i, j = i + 1, j - 1 {
                if s[i] != s[j] {
                    flag2 = false
                    break
                }
            }
            return flag1 || flag2
        }
    }
    return true
}

