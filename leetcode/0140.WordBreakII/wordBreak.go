package lt140

import "strings"

/*difficult   lt139 I
140. 单词拆分 II
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

说明：

分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
*/
/*本题ＣＶ
dfs?  回溯＋剪枝
自顶向下的记忆化搜索 将不可拆分的情况剪枝
s 如果某个前缀是单词列表中单词，则拆分
通过记忆化降低时间复杂度

*/
func wordBreak(s string, wordDict []string) []string {
	//wordDict转为map
	wordSet := map[string]struct{}{}
	for _,word := range wordDict {
		wordSet[word] = struct{}{}
	}
	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(idx int) [][]string 
	backtrack = func(idx int) [][]string {
		if dp[idx] != nil {
			return dp[idx]
		}
		wordList := [][]string{}
		for i:=idx+1;i<n;i++{
			word := s[idx:i]  //前缀
			if _,has := wordSet[word]; has {
				for _,nextWords := range backtrack(i) {
					wordList = append(wordList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[idx:]
		if _,has := wordSet[word];has{
			wordList = append(wordList, []string{word})
		}
		dp[idx] = wordList
		return wordList
	}
	res := []string{}
	for _,words := range backtrack(0) {
		res = append(res, strings.Join(words, " "))
	}
	return res
}