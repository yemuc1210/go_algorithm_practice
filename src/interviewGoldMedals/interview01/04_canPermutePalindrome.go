package interview01
/*
面试题 01.04. 回文排列

给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。



由题意可知要使得原字符串能构成一个回文串，则最多只能有一个字母出现奇数次。
排序后遍历，对出现的每个字母统计其出现次数；若出现奇数次的字母个数大于1直接跳出循环。
循环结束后 注意要统计字符串最后出现的字母的出现次数。

*/
func canPermutePalindrome(s string) bool {
    strMap := make(map[rune]int, 26)
    for _, v := range s {
        strMap[v] = strMap[v] + 1
    }
    odd := 0
    for _, v := range strMap {
        if v % 2 != 0 {
            odd++
        }
    }
    return odd == 0 || odd == 1

}