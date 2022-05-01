package tencent50
/*
557. 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
示例：
	输入："Let's take LeetCode contest"
	输出："s'teL ekat edoCteeL tsetnoc"
 
提示：
在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

*/
func reverseWords(s string) string {
	if len(s)==0{
		return ""
	}
	//处理前置空格
	index := 0 
	for index<len(s) && s[index]==' ' {
			index ++
	}
	words := []string{}
	var res string
	for index<len(s){
		if s[index] != ' '{
				//则取单词
				var tmp []byte
				for index<len(s) && s[index] !=' ' {
						tmp = append(tmp, s[index])
					
						index ++
				}
				words = append(words, string(tmp))
		}
		for index<len(s) && s[index] ==' ' {
				index++
		}
	}
	//逆序输出
	for i:=0;i<len(words);i++{
		if i == len(words)-1 {
				res+=reverseString557([]byte(words[i]))
		}else{
				res += reverseWords( words[i])+" "
		}
	}
	return res
}
//调用reverseString
func reverseString557(s []byte) string {
	n := len(s)
	for i:=0; i< n/2; i++{
			s[i],s[n-i-1] = s[n-i-1],s[i]
	}
	return string(s)
}