package tencent50
/*
217. 存在重复元素
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
*/
/*
可能都不相同 
思路1：map

思路2：摩尔投票
	如果都不相同，voteCnt=0
	如果存在重复元素 votCnt!=0 
	错误 【1，2，3，4】
		最终 votCnt==1
*/
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