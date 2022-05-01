package offer50

/*
剑指 Offer 50. 第一个只出现一次的字符
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

排序
哈希表  两次遍历 第一次得出次数  第二次得出答案
*/
func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _,ch := range s{
		cnt[ch - 'a'] ++
	}
	for i,ch := range cnt {
		if cnt[ch-'a'] == 1{
			return s[i]
		}
	}
	return ' '
}