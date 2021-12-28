package main

 

/*
472. 连接词
给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。

连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。

示例 1：

输入：words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
输出：["catsdogcats","dogcatsdog","ratcatdogcat"]
解释："catsdogcats" 由 "cats", "dog" 和 "cats" 组成;
     "dogcatsdog" 由 "dog", "cats" 和 "dog" 组成;
     "ratcatdogcat" 由 "rat", "cat", "dog" 和 "cat" 组成。

*/
/*
长度上有关系
或者直接得到所有可能的组合词，然后查找是否存在在数组中，得到连接词
首先根据长度限制：即最长的单词的长度，得到组合词长度的上限
//因而可能不止两个单词 所以需要将组合的词插入道原来的words元组
*/
func findAllConcatenatedWordsInADict(words []string) (ret []string) {
	set := map[string]bool{}

	for _, word := range words {
		set[word] = true
	}

	var canBreak func(s string, set map[string]bool) bool
	canBreak = func(s string, set map[string]bool) bool {
		if len(set) == 0 || len(s) == 0 {
			return false
		}
		//dp数组，dp[i]表示word从【0，i)的字符串是否可以分解成一个更小的【wiord】 
		//dp[n]表示可以完全分解
		dp := make([]bool, len(s)+1)
        dp[0] = true
		for i := 1; i < len(s)+1; i++ {
			for j := 0; j < i; j++ {
				if !dp[j] {
					continue
				}
				//dp[j]=true   [0,j)是更小的字符串  
				if set[s[j:i]] {   //字符串s[j:i]是一个有效的小字符串
					//所以[0,i)是一个可以分解的字符串
					dp[i] = true
					break
				}
			}
		}
		return dp[len(s)]  
	}

	for _, word := range words {
		if "" == word {
			continue
		}
		delete(set, word)

		if canBreak(word, set) {
			ret = append(ret, word)
		}
		set[word] = true
	}
	return
}

func main(){
	words := []string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}
	findAllConcatenatedWordsInADict(words)
}