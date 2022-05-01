package lt46
/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。
你可以 按任意顺序 返回答案。
示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permute(nums []int) [][]int {
	//递归
	if len(nums)==1 || len(nums)==0{
		return [][]int{nums}
	}
	res := [][]int{}
	used :=make([]bool,len(nums))
	p:=[]int{}
	solve(nums,&res,&used,0,p)
	return res
}
func solve(nums []int,res *[][]int,used *[]bool,idx int,p []int){
	if idx==len(nums){
		tmp:=make([]int,len(p))
		copy(tmp,p)
		*res = append(*res, tmp)
		return 
	}
	for i:=0;i<len(nums);i++{
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			solve(nums,res,used,idx+1,p)
			p=p[:len(p)-1]
			(*used)[i]=false
		}
	}
}

func permute1(nums []int) [][]int{
	res:=[][]int{}
	visited := map[int]bool{}   //记录排列是否访问过

	var dfs func(path []int)
	dfs = func(path []int){
		if len(path)==len(nums){// 长度相同，出现答案
			tmp := make([]int,len(path))
			copy(tmp,path)
			res = append(res, tmp)
			return
		}
		for _,n := range nums{
			if visited[n]{
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]    //返回之前的状态
			visited[n]=false
		}
	}
	dfs([]int{})
	return res
}