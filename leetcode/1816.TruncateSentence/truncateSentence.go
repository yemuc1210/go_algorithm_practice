package lt1816
/*
1816. 截断句子
句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。每个单词仅由大小写英文字母组成（不含标点符号）。

例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。
给你一个句子 s​​​​​​ 和一个整数 k​​​​​​ ，请你将 s​​ 截断 ​，​​​使截断后的句子仅含 前 k​​​​​​ 个单词。返回 截断 s​​​​​​ 后得到的句子。

 
*/
func truncateSentence(s string, k int) string {
	var sentences []string

	var idx int
	var n = len(s)
	for idx < n {
		word := []byte{}
		for idx<n && s[idx] != ' ' {
			word = append(word, s[idx])
			idx ++
		}
		//跳过中间的空格
		for idx<n && s[idx] == ' ' {
			idx ++
		}
		//保存单词
		sentences = append(sentences, string(word))
	}
	//返回前k个
	var res string
	if len(sentences) <= k {
		//返回全部的sentences内容
		for i:=0;i<len(sentences);i++{
			if i== len(sentences)-1 {
				res += sentences[i]
			}else{
				res += sentences[i]+" "	
			}
		}
		return res
	}else{
		for i:=0;i<k;i++{
			if i==k-1{
				res += sentences[i]
			}else{
				res += sentences[i]+" "
			}
		}
	}
	return res
}