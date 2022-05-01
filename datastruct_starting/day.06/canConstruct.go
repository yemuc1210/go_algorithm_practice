package day06
/*
给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，
判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。
如果可以构成，返回 true ；否则返回 false。

(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。
	杂志字符串中的每个字符只能在赎金信字符串中使用一次。)


示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
*/

/*
统计字符及个数
*/
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine)<len(ransomNote){
		return false      //字符数都不够，肯定不行
	}
	charRansom := make([]int,26) // 统计26个字符的频次
	
	for _,v := range ransomNote{
		charRansom[v-'a']++
	}
	for i:=0;i<len(magazine);i++{
		if charRansom[magazine[i] - 'a'] > 0 { //说明这个字符是需要的
			charRansom[magazine[i]-'a']--
		}
	}
	for _,v :=range charRansom{
		if v>0{
			return false
		}
	}
	return true
}