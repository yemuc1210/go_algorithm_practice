package offerS85
/* 回溯  lt22
剑指 Offer II 085. 生成匹配的括号
正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

//dfs 有左必有右
func generateParenthesis(n int) []string {
	if n==0 {
		return []string{}
	}
	res := []string{}
	dfs(&res, n,n, "")

	return res
}
//ridx lidx 理解为 remain
//
func dfs(res *[]string, ridx,lidx int, s string){
	if lidx==0 && ridx==0 {
		*res = append(*res, s)
		return 
	}

	if lidx > 0{
		dfs(res, ridx,lidx-1,s+"(")
	}
	if ridx > 0 && lidx<ridx{
		dfs(res,ridx-1,lidx,s+")")
	}
}
