package lt205
/*  lt290 模式匹配
205. 同构字符串
给定两个字符串 s 和 t，判断它们是否是同构的。
如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:
输入：s = "egg", t = "add"
输出：true
示例 2：
输入：s = "foo", t = "bar"
输出：false
示例 3：
输入：s = "paper", t = "title"
输出：true。
*/
/*
首先，两个字符串得相等
然后：寻求映射规则，
	想直接得到明确的映射函数比较困难
使用两个map记录映射关系  一个map记录s-t 另一个记录t-s

*/
func isIsomorphic(s string, t string) bool {
	strList := []byte(t) 
	patternByte := []byte(s)
	if (s==""&&t!="") ||(len(strList)!=len(patternByte)) {
		return false
	}

	pMap,sMap := map[byte]byte{}, map[byte]byte{}

	for idx,b := range patternByte{
		if _,ok := pMap[b]; !ok {
			if _,ok := sMap[strList[idx]]; !ok {
				pMap[b] = strList[idx]
				sMap[strList[idx]] = b //双向映射
			} else {
				if sMap[strList[idx]] !=b {
					return false  //映射对不上
				}
			}
		}else {
			if pMap[b] != strList[idx] {
				return false
			}
		}
	}
	return true
}