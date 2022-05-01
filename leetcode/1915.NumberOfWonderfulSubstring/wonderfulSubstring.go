package main
// 最美字符串：至多一个字母出现奇数次
// word: 前十个小写字母组成 'a'-'j'

// 求最美子字符串的个数
// 子字符串！=子序列

func wonderfulSubstrings(word string) int64 {
    return wonderfulSubstringsZero(word) + wonderfulSubstringsOne(word)
}

// 字符串每个字符都出现偶数次
func wonderfulSubstringsZero(word string) int64 {
    m := map[int]int{0:1}
    cur := 0
    res := 0
    for _, ch := range word{
        cur ^= 1 << (ch - 'a')
        res += m[cur]
        m[cur] ++
    }
    return int64(res)
}

// 字符串存在一个字符出现奇数次
func wonderfulSubstringsOne(word string) int64 {
    m := map[int]int{0:1}
    cur := 0
    res := 0
    for _, ch := range word{
        cur ^= 1 << (ch - 'a')
        for i := 0; i < 10; i ++{
            res += m[1 << i ^ cur]
        }
        m[cur] ++
    }
    return int64(res)
}
