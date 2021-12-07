package offerS15
/*
剑指 Offer II 015. 字符串中的所有变位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。
不考虑答案输出的顺序。

变位词 指字母相同，但排列不同的字符串。
*/
/*
这道题是一道考“滑动窗口”的题目。和第 3 题，第 76 题，第 567 题类似的。
解法也是用 freq[256] 记录每个字符的出现的频次次数。
滑动窗口左边界往右滑动的时候，划过去的元素释放次数(即次数 ++)，
滑动窗口右边界往右滑动的时候，划过去的元素消耗次数(即次数 --)。
右边界和左边界相差 len(p) 的时候，需要判断每个元素是否都用过一遍了。

具体做法是每经过一个符合规范的元素，count 就 --，count 初始值是 len(p)，
当每个元素都符合规范的时候，右边界和左边界相差 len(p) 的时候，count 也会等于 0 。
当区间内有不符合规范的元素(freq < 0 或者是不存在的元素)，
那么当区间达到 len(p) 的时候，count 无法减少到 0，
区间右移动的时候，左边界又会开始 count ++，
只有当左边界移出了这些不合规范的元素以后，才可能出现 count = 0 的情况，
即找到 Anagrams 的情况。

*/
func findAnagrams(s string, p string) []int {  //s是长  p短 s中找p
	if len(s) < len(p) {
		return []int{}
	}
	res := make([]int,0)
	left,right := 0,0
	length := len(s)
	aim := make(map[byte]int)
	now := make(map[byte]int)
	nowNum := 0

	for i := range p{
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