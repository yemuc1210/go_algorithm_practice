package offerS83

/*
剑指 Offer II 083. 没有重复元素集合的全排列
给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
*/
//dfs
func permute(nums []int) [][]int {
	res := [][]int{}
	visited := make(map[int]bool)  //数组不含重复元素

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {//长度相同，出现答案
			res = append(res, append([]int(nil),path...))
			return 
		}

		for _,n := range nums{
			if visited[n] {
				continue //被访问过，跳过
			}
			path = append(path, n)  //否则，未访问过
			visited[n] = true  //访问之
			dfs(path)   //继续递归
			path = path[:len(path)-1]  //返回之前的状态
			visited[n] = false  //不访问n的情况再次递归
		}
	}
	dfs([]int{})
	return res
}	