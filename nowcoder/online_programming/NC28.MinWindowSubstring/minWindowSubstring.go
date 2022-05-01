package NC28

/*
76. 最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/

/*
在窗口滑动的过程中不断的包含字符串 T，直到完全包含字符串 T 的字符以后，
记下左右窗口的位置和窗口大小。
每次都不断更新这个符合条件的窗口和窗口大小的最小值。
最后输出结果即可。
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