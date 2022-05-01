package offerS81
/*  lt39
剑指 Offer II 081. 允许重复选择元素的组合
给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，
找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。

candidates 中的数字可以无限制重复被选取。
如果至少一个所选数字数量不同，则两种组合是唯一的。 

对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
*/
func combinationSum(candidates []int, target int) [][]int {
	combine := []int{}//已经组合的列表
	res := [][]int{}
	var dfs func(target int,idx int)   //剩余组合目标target  当前位置idx
	dfs = func(target,idx int){
		if idx == len(candidates) {
			return 
		}
		if target==0 {
			res = append(res, append([]int(nil),combine...))
			return
		}
		//跳过当前idx元素
		dfs(target,idx+1)
		if target - candidates[idx] >= 0{
			combine = append(combine, candidates[idx])
			dfs(target-candidates[idx], idx)   //因为元素可以无限选取，所以idx
			combine = combine[:len(combine)-1]  //恢复  
		}
	}
	dfs(target,0)
	return res
}