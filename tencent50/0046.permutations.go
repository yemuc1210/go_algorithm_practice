package tencent50
/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/
/*
dfs   参数idx,
全排列：1
*/
func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}  //记录排列是否访问过
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			//得到一组结果
			tmp := make([]int,len(path))
			copy(tmp,path)
			res = append(res, tmp)
			return 
		}
		for _,n := range nums{
			if visited[n] {
				continue
			}
			//否则，选择未访问过的
			path = append(path, n)
			visited[n] = true
			dfs(path)
			//返回之前的状态
			path = path[:len(path)-1]
			visited[n]=false    //也就是回溯的感觉
		}
	}
	dfs([]int{})
	return res
}