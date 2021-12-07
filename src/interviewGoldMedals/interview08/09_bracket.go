package interview08
/*
v面试题 08.09. 括号
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。

说明：解集不能包含重复的子集。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/
func generateParenthesis(n int) []string {
	if n <= 0{
		return []string{}
	}

	res := []string{}
	var dfs func(paths string, left,right int)
	dfs = func(paths string, left,right int) {
		if left>n || right > left {
			return 
		}
		if len(paths) == n*2 {
			//括号成对出现
			res = append(res, paths)
		}

		dfs(paths+"(", left+1,right)
		dfs(paths+")", left, right+1)
	}

	dfs("",0,0)
	return res
}