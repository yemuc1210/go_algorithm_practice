package leetcode

/*
数字 n 代表生成括号的对数，   ()
请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

//dfs吧
func generateParenthesis(n int) []string {
	if n==0{
		return []string{}
	}
	res := []string{}
	dfs(n,n,"",&res)
	return res
}

func dfs(lindex,rindex int,s string,res *[]string){
	if lindex==0&&rindex==0{
		*res = append(*res, s)
		return
	}
	if lindex>0{
		dfs(lindex-1,rindex,s+"(",res)
	}
	if rindex>0 && lindex<rindex{
		dfs(lindex,rindex-1,s+")",res)
	}
}