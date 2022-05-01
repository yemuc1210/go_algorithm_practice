package lt1941
/*
1941. 检查是否所有字符出现次数相同
给你一个字符串 s ，如果 s 是一个 好 字符串，请你返回 true ，否则请返回 false 。
如果 s 中出现过的 所有 字符的出现次数 相同 ，那么我们称字符串 s 是 好 字符串。
*/
func areOccurrencesEqual(s string) bool {
	if len(s) == 0 {
		return false
	}
	//map
	m := make(map[rune]int)
	for _,ch := range s{
		m[ch] ++
	}
	cntArr := []int{}
	for _,cnt := range m {
		cntArr = append(cntArr, cnt)
	}
	cnt := cntArr[0]
	for i:=1;i<len(cntArr);i++{
		if cnt!=cntArr[i] {
			return false
		}
	}
	return true
}