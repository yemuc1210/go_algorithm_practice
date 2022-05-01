package lt997
/*
编号1~n
法官不相信任何人；每个人都信任法官；法官只有一位
trust[i]=[a,b]表示a信任b

存在a不信任任何人，但是a不是法官
*/
func findJudge(n int, trust [][]int) int {
	if n==1{
		return 1
	}
	relationship := make([][]int, n)
	for i:=0;i<n;i++{
		relationship[i] = make([]int, 2)
	}
	for _,p := range trust {
		out,in := p[0],p[1]   //out有出边(信任某人)      in有入边(被信任)
		relationship[out-1][0] +=1    
		relationship[in-1][1] += 1
	}
	//遍历
	for i:=0;i<n;i++{
		if relationship[i][0] == 0 && relationship[i][1]==n-1{
			return i + 1  //序号1~n
		}
	}
	return -1

}