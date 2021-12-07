package offerS79
/*递归  回溯
剑指 Offer II 079. 所有子集
给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
func subsets(nums []int) [][]int {
	res := [][]int{}
	//子集不重复  元素互不相同
	set := []int{}
	var dfs func(int)
	dfs = func(cur int){
		if cur == len(nums)-1{//得到一个子集
			res = append(res, append([]int(nil),set...))
			return 
		} 
		set = append(set, nums[cur])
		dfs(cur+1)  //下一个idx
		set = set[:len(set)-1]//刚加入的退出  回溯
		dfs(cur+1)
	}
	dfs(0)
	return res
}
