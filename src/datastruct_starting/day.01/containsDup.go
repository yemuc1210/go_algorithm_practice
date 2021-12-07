/*
数据结构入门第一天
给定一个整数数组，判断是否存在重复元素。
如果存在一值在数组中出现至少两次，函数返回 true 。
如果数组中每个元素都不相同，则返回 false 。

示例 1:
输入: [1,2,3,1]
输出: true
*/
package dsday01
func containsDuplicate(nums []int) bool {
	//方法一、使用map,并计数
	//效率26ms  10.37%       7.7MB    65.60%
	mp := map[int]int{}   //key:元素值   val:次数
	for _,key := range nums {
		mp[key] ++
		if mp[key] > 1{
			return true
		}
	}
	return false
}
//20 ms 94.38%        8MB  38.97%
func containsDuplicate1(nums []int) bool {
    set := map[int]struct{}{}
    for _, v := range nums {
        if _, has := set[v]; has {
            return true
        }
        set[v] = struct{}{}
    }
    return false
}
// func containsDuplicate(nums []int) bool {
// 	type void struct{}
// 	var v void

// 	set := make(map[int]void)
// 	for _, item := range nums {
// 		set[item] = v
// 	}
// 	isEqual := len(nums) == len(set)
// 	return !isEqual

// }
