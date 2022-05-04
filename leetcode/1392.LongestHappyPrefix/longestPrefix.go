package main

func longestPrefix(s string) string {
    N := len(s)
    first, last := s[0], s[N-1]
    // j 从大的开始往前移  j从小的往后移动  贪心思想 一旦匹配成功了,就是最长的答案了,直接return
    for i,j := 0, N-1; i < N && j >= 0; i, j = i+1, j-1 {
        if s[j] == last && s[i] == first && s[:j+1]==s[i:] && s[:j+1] != s {
            return s[:j+1]
        }
    }
    return ""
}

