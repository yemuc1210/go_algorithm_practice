package leetcode 


/*
给定仅包含2-9的数字串，返回所有可能的字母组合
使用DFS深度遍历

*/
func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	var res []string

	letterMap := []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	dfs(0,letterMap,&digits,&res,"")	
	return res
}

func dfs(cur int,letter_map []string,digits *string,res *[]string,s string){
	if cur ==len(*digits){
		*res = append(*res,s)
		return
	}
	num := (*digits)[cur]
	letter := letter_map[num-'0']
	for i:=0;i<len(letter);i++ {
		dfs(cur+1,letter_map,digits,res,s+string(letter[i]))
	}

}