package tencent50
/* 简
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
/*
一个比较费空间的思路：
	记录每个前缀的 出现次数   出现次数==len(strs)
构建树  
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//找出长度最长的字符串
	maxn,maxstr := 0,""
	for _,str := range strs{
		// n := len(str)
		if len(str) > maxn {
			maxn = len(str)
			maxstr = str
		}
	}
	//根据最长的字符串构建树，因为最长前缀必然在其中
	bytes := []byte(maxstr)  //转换成字符数组
	minCut := maxn  //记录最长公共前缀的最后一位
	for _,str := range strs {
		if str != maxstr {
			cntTmp := 0
			byteTmp := []byte(str)
			for i,ch := range byteTmp {
				if ch == bytes[i] {
					cntTmp++  //仍然是前缀
				}else{
					if i==0 {
						//最差情况 第一个就不同 不存在公共前缀
						return ""
					}
					break
				}
			}
			if cntTmp < minCut {  //遍历完当前字符串，更新一下公共前缀 找出最小值
				minCut = cntTmp
			}
		}
	}
	return maxstr[:minCut]
}