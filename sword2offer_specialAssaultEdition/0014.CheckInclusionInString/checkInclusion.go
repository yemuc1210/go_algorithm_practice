package offerS14
/*   leetcode 567
剑指 Offer II 014. 字符串中的变位词
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。

换句话说，第一个字符串的排列之一是第二个字符串的 子串 。

求s1的排列组合，然后判断是否为字串

直接根据字符数量判断
滑动窗口：两个字符每个字符的数量均相等，才可以是变位词。
做法：记n=len(s1) 遍历s2中每个长度为n的子串，判断子串和s1中字符个数
*/
func checkInclusion(s1 string, s2 string) bool {
	var freq [256]int
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}
	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--  
		right++  //边界右移
		if count == 0 {
			return true
		}
		if right-left == len(s1) {   //区间长度相等
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}