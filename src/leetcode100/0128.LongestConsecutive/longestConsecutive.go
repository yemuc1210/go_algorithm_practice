package lt128

/*lt128  中 图
剑指 Offer II 119. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
/*
未排序
暴力思路：每个数都存在 `map` 中，先删去 `map` 中没有前一个数 `nums[i]-1` 也没有后一个数 `nums[i]+1` 的数 `nums[i]`，
这种数前后都不连续。然后在 `map` 中找到前一个数 `nums[i]-1` 不存在，
但是后一个数 `nums[i]+1` 存在的数，这种数是连续序列的起点，那么不断的往后搜，直到序列“断”了。最后输出最长序列的长度。
*/
func longestConsecutive(nums []int) int {
	if len(nums)==0{
		return 0
	}
	numMap, length, tmp, lcs := map[int]bool{}, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = true
	}
	for key := range numMap {
		if !numMap[key-1] && !numMap[key+1] {
			delete(numMap, key)
		}
	}
	if len(numMap) == 0 {
		return 1
	}
	for key := range numMap {
		if !numMap[key-1] && numMap[key+1] {
			length, tmp = 1, key+1
			for numMap[tmp] {
				length++
				tmp++
			}
			lcs = max(lcs, length)
		}
	}
	return max(lcs, length)
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
