package offerS17
/*

剑指 Offer II 017. 含有所有字符的最短字符串
给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。
如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。

如果 s 中存在多个符合条件的子字符串，返回任意一个。


注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。

 
*/
func minWindow(s string, t string) string {
	if s == "" || t == ""{
		return ""
	}
	//求t的所有字符
	var tFreq,sFreq [256]int
	left,right := 0,-1
	minWin := len(s)+1
	count := 0
	res := ""
	//需要返回字串，所以需要记录最小字串的下标
	minLeftIds,minRightIdx := -1,-1

	for i:=0;i<len(t);i++ {
		tFreq[t[i]-'a'] ++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			//+
			sFreq[s[right+1] - 'a'] ++
			//判断  还没完全包含s[right+1]这个字符
			if sFreq[s[right+1] - 'a'] <= tFreq[s[right+1] - 'a'] {
				count ++
			}
			right ++
		}else {  //count == len(t)  此时包含了所有字符
			if right-left+1 < minWin && count == len(t) {  
				//更新最小值
				minWin = right-left+1
				minLeftIds,minRightIdx = left,right
			}
			//左移
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count --
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if minLeftIds != -1 {
		for i := minLeftIds;i<minRightIdx+1;i++{
			res += string(s[i])
		}
	}
	return res
}