package lt438


/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指字母相同，但排列不同的字符串。

异位词判断用map
 
滑动窗口题
*/
func findAnagrams(s string, p string) []int {
	//窗口值是p的大小
	if len(s) < len(p) {
		return []int{}
	}
	res := make([]int,0)
	left,right := 0,0
	length := len(s)
	aim := make(map[byte]int)
	now := make(map[byte]int)
	nowNum := 0

	for i,_ := range p{
		aim[p[i]] ++
	}
	aimNum := len(aim) //字符种类

	for right < length {
		for nowNum != aimNum && right<=length-1 {
			//移动右指针，窗口扩大
			if aim[s[right]] !=0 && aim[s[right]] > now[s[right]] {
				//判断字符是否符合要求，且该种类字符数量尚未满足
				now[s[right]] ++
				if now[s[right]] == aim[s[right]]{
					nowNum ++
				}
			}else if aim[s[right]] != 0 && aim[s[right]] <= now[s[right]] {//如果该种类字符数量满足了，则只加个数，不加种类数
				now[s[right]] += 1

			}
			right += 1
		}
		for nowNum == aimNum && (right-left) <= length {//移动左指针
			if right-left == len(p) {//当长度相等，且种类数相等时，一定满足条件
				res = append(res, left)
			}
			if aim[s[left]] != 0 {
				now[s[left]] -= 1
				if now[s[left]] < aim[s[left]] {
					nowNum -= 1
				}
			}

			left += 1
		}
	}
	return res
}