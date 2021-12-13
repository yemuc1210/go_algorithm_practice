package NC26

/**
  * 给出n对括号，生成所有合法组合   lt22
  * @param n int整型 
  * @return string字符串一维数组
*/
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
	}//先左后右
	if rindex>0 && lindex<rindex{
		dfs(lindex,rindex-1,s+")",res)
	}
}