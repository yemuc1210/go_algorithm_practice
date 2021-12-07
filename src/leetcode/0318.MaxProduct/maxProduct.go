package lt318


/*每日一题 11167
318. 最大单词长度乘积
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，
并且这两个单词  不含有公共字母  。你可以认为每个单词只包含小写字母。
如果不存在这样的两个单词，返回 0。
*/

/*
关键：两个字符不含有共同字母
长度 排序可做  不可行
动态规划似乎也可
*/
func maxProduct(words []string) int {
	/**
	全是小写字母, 可以用一个32为整数表示一个word中出现的字母, 
	hash[i]存放第i个单词出现过的字母, a对应32位整数的最后一位,
	b对应整数的倒数第二位, 依次类推. 时间复杂度O(N^2)
	判断两两单词按位与的结果, 如果结果为0且长度积大于最大积则更新
	**/
	n := len(words)
	hash := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		for _,v := range words[i] {
			hash[i] |= 1 << (v-'a')
		}
	}
	
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if (hash[i] & hash[j]) == 0{
				res = max(len(words[i]) * len(words[j]), res)
			}
		}
	}
	return res
}

func max(a,b int) int{
   if a>b{
	   return a
   }
   return b
}