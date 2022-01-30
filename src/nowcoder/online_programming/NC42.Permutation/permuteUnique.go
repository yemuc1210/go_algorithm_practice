package NC42

import "sort"

/**
 * 有重复项数字的排列 offerII84  lt47
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func permuteUnique( nums []int ) [][]int {
    // 排序去重
	sort.Ints(nums)
	n,permute := len(nums), []int{}
	visited := make([]bool, n)

	res := [][]int{}
	//dfs
	var dfs func(int)
	dfs = func(idx int) {
		if idx==n { //得到一组排列
			res = append(res, append([]int(nil), permute...))
			return 
		}
		for i,v := range nums {
			if visited[i] || i>0 && !visited[i-1] && v==nums[i-1] {
				continue //访问过  值访问过
			}
			permute = append(permute, v)
			visited[i] = true

			dfs(idx+1) 
			visited[i]=false//回溯
			permute = permute[:len(permute)-1]
		}
	}
	dfs(0)
	return res
}