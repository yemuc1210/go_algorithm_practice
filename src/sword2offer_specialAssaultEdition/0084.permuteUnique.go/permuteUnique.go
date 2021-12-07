package offerS84

import "sort"

/*
剑指 Offer II 084. 含有重复元素集合的全排列
给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
*/
func permuteUnique(nums []int) [][]int {
	//先排序
	sort.Ints(nums)
	n,permute := len(nums),[]int{}
	visited := make([]bool,n)
	
	res := [][]int{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			res = append(res, append([]int(nil), permute...))
			return 
		}
		for i,v := range nums{
			if visited[i] || i>0 && !visited[i-1] && v== nums[i-1] {
				//可能包含重复值v==nums[i-1]  跳过重复值
				continue  //边界  及   不可能条件
			}
			permute = append(permute, v)
			visited[i] = true

			dfs(idx+1)
			visited[i] = false
			permute = permute[:len(permute)-1]
		}
	}
	dfs(0)
	return res
}