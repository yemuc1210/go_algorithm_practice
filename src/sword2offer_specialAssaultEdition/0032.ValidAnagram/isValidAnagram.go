package offerS32
/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

 

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
数组或哈希表累加 s 中每个字符出现的次数，再减去 t 中对应的每个字符出现的次数。
遍历结束后，若字符中出现次数不为 0 的情况，返回 false，否则返回 true。


*/
func isAnagram(s string, t string) bool {
    if len(s) != len(t) || s == t {
        return false
    }
    var chars [26]int
    for i := 0; i < len(s); i++ {
        chars[s[i]-'a']++
        chars[t[i]-'a']--
    }
    for _, c := range chars {
        if c != 0 {
            return false
        }
    }
    return true
}

