package lt219

/*
219. 存在重复元素 II
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
*/
/*
定义一个新的数据结构 pair{originIdx,val}
然后根据val进行排序

本题是判断是否存在
	设置窗口尺寸k  要判断当前值是否符合条件的另一个数，在这个窗口内寻找
	复杂度是O(n^2)  时间复杂度略大

map
	map尺寸为k 每次判断map中是否存在当前元素，若存在则return true
	若长度超过K,则删除i-k那个元素
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1{
		return false
	}
	if k<=0{
		return false
	}
	m := make(map[int]bool, len(nums))
	for i,n := range nums{
		if _,found:= m[n]; found {
			return true
		}
		m[n] = true
		if len(m) == k+1 {
			delete(m, nums[i-k])
		}
	}
	return false
}