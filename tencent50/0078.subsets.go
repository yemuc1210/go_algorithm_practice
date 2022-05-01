package tencent50
/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。
返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

/*
dfs  终止条件 没想好
	空
	一个元素
	两个元素
	三个元素
	。。。

用1 表示字符在子集  0表示不在
则对应二进制数0-2^n-1
利用mask∈[0,2^n−1] 枚举
根据mask在集合中取值
*/
func subsets(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	for mask:=0;mask < 1<<n ; mask++{
		set := []int{}
		for i,v := range nums{ //根据mask取集合中元素组成子集
			if mask>>i&1 > 0 {   //第i位是否为1
				set = append(set, v)  //若是，则取第i个元素
			}
		}
		//得到一组子集
		res = append(res, append([]int(nil),set...))
	}
	return res
}