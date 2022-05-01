package main

/*
给你一个整数数组 nums ，请你找出 nums 子集 按位或 可能得到的 最大值 ，
并返回按位或能得到最大值的 不同非空子集的数目 。

如果数组 a 可以由数组 b 删除一些元素（或不删除）得到，
则认为数组 a 是数组 b 的一个 子集 。如果选中的元素下标位置不一样，则认为两个子集 不同 。

对数组 a 执行 按位或 ，结果等于 a[0] OR a[1] OR ... OR a[a.length - 1]（下标从 0 开始）。

非空子集数量 2^n-1
用长度为n比特的整数表示不同子集。为1则表示选择这个元素
求出每个子集按位或的值
*/

// 1. 先求所有子集，然后得到其中或运算最大的
func countMaxOrSubsets(nums []int) int {
	maxOr := 0 
	res := 0

	for i:=1;i<1<<len(nums);i++{
		or := 0 
		for j,num := range nums {
			// 根据i 判断哪些数字需要取用
			if i>>j&1 == 1 {
				// 第j个需要
				or |= num
			} 
		}
		if or > maxOr {
			maxOr = or
			res = 1
		}else if or == maxOr {
			res ++
		}
	}
	return res
}

func main(){

}