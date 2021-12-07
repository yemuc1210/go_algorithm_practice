package lt151
/*中  剑指offer 58
151. 翻转字符串里的单词
给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
说明：
	输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
	翻转后单词间应当仅用一个空格分隔。
	翻转后的字符串中不应包含额外的空格
提示：
	1 <= s.length <= 104
	s 包含英文大小写字母、数字和空格 ' '
	s 中 至少存在一个 单词。
*/
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	//处理前置空格
	idx,n := 0,len(s)
	for idx<n && s[idx] ==' '{idx++}
	words := []string{}
	var res string
	for idx < n {
		if s[idx]!=' ' {
			//判断是否是合法字符 本题其实不需要
			// (s[idx]>='a'&&s[idx]<='z') && (s[idx]>='A'&&s[idx]<='Z') &&(s[idx]>='0' && s[idx]<='9')
			var tmp []byte
			for idx<n &&s[idx]!=' ' &&(s[idx]>='a'&&s[idx]<='z') && (s[idx]>='A'&&s[idx]<='Z') &&(s[idx]>='0' && s[idx]<='9'){
				tmp = append(tmp, s[idx])
				idx ++
			} 
			words = append(words, string(tmp))
		}
		//处理单词间空格
		for idx<n && s[idx]==' '{idx++}
	}
	//开始所有的单词
	for i:=len(words)-1;i>=0;i--{
		if i==0 {
			res += words[i]
		}else{
			res += words[i]+" "
		}
	}
	return res
}