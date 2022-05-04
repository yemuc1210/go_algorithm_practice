package lt524
/*
524. 通过删除字母匹配到字典里最长单词
给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。

如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。


匹配
双指针
根据题意，我们需要解决两个问题：
	如何判断 dictionary 中的字符串 t 是否可以通过删除 s 中的某些字符得到；
	如何找到长度最长且字典序最小的字符串。

这样，我们初始化两个指针 i 和 j，分别指向 t 和 s 的初始位置。
每次贪心地匹配，匹配成功则 i 和 j 同时右移，匹配 t 的下一个位置，
匹配失败则 j 右移，i 不变，尝试用 s 的下一个字符匹配 t。

最终如果 i 移动到 t 的末尾，则说明 t 是 s 的子序列。

第 2 个问题可以通过遍历 dictionary 中的字符串，并维护当前长度最长且字典序最小的字符串来找到。


*/
func findLongestWord(s string, dictionary []string) (ans string) {
    for _, t := range dictionary {
        i := 0
        for j := range s {
            if s[j] == t[i] {
                i++
                if i == len(t) {
                    if len(t) > len(ans) || len(t) == len(ans) && t < ans {
                        ans = t
                    }
                    break
                }
            }
        }
    }
    return
}

