package lt500

import "unicode"

/*
500. 键盘行
给你一个字符串数组 words ，
只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示

美式键盘 中：

第一行由字符 "qwertyuiop" 组成。
第二行由字符 "asdfghjkl" 组成。
第三行由字符 "zxcvbnm" 组成。
*/
func findWords(words []string) []string {
	const chIdx string = "23321222122233111121131313"  //26字母的位置line
	res := []string{}
next:
	for _,word := range words{
		oneIdx := chIdx[ unicode.ToLower(rune(word[0]))-'a']
		for i:=1;i<len(word);i++{
			if chIdx[ unicode.ToLower(rune(word[i]))-'a' ] != oneIdx {
				// break  //不行
				continue next
			}
		}
		//如果顺利执行
		res = append(res, word)
	}
	return res
}