package offerS114

import "strings"

/*lt 269 困难 不会做
剑指 Offer II 114. 外星文字典
现有一种使用英语字母的外星文语言，这门语言的字母顺序与英语顺序不同。

给定一个字符串列表 words ，作为这门语言的词典，
	words 中的字符串已经 按这门新语言的字母顺序进行了排序 。

请你根据该词典还原出此语言中已知的字母顺序，并 按字母递增顺序 排列。
若不存在合法字母顺序，返回 "" 。
若存在多种可能的合法字母顺序，返回其中 任意一种 顺序即可。

字符串 s 字典顺序小于 字符串 t 有两种情况：

在第一个不同字母处，如果 s 中的字母在这门外星语言的字母顺序中位于 t 中字母之前，那么 s 的字典顺序小于 t 。
如果前面 min(s.length, t.length) 字母都相同，那么 s.length < t.length 时，s 的字典顺序也小于 t 。
示例 1：

输入：words = ["wrt","wrf","er","ett","rftt"]
输出："wertf"
*/
/*
字符串已经按字母顺序排序  首字母排序-第二个-第三个....


*/
func alienOrder(words []string) string {
	graph := make(map[rune]map[rune]bool)
	inDegrees := make(map[rune]int)
	for _, word := range words {
		for _, c := range word {
			if _, has := graph[c]; !has {
				graph[c] = make(map[rune]bool)
				inDegrees[c] = 0
			}
		}
	}
	n := len(words)
	for i := 1; i < n; i++ {
		var j int
		length := min(len(words[i-1]), len(words[i]))
		for ; j < length; j++ {
			ch1 := rune(words[i-1][j])
			ch2 := rune(words[i][j])
			if ch1 != ch2 {
				if _, has := graph[ch1][ch2]; !has {
					graph[ch1][ch2] = true
					inDegrees[ch2]++
				}
				break
			}
		}
		if j == length && len(words[i-1]) > len(words[i]) {
			return ""
		}
	}

	var ret strings.Builder
	var queue []rune
	for i, d := range inDegrees {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		ret.WriteRune(ch)
		for c := range graph[ch] {
			inDegrees[c]--
			if inDegrees[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	if ret.Len() != len(inDegrees) {
		return ""
	}
	return ret.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
 