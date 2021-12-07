package offer10

import (
	"fmt"
	"testing"
)

type question10 struct {
	par10
	ans10
}
type par10 struct {
	n int
}

type ans10 struct {
	output int
}

func Test_FrogStair(t *testing.T) {
	qs := []question10{
		{
			par10{2}, ans10{2},
		},
		{
			par10{7}, ans10{21},
		},
		{
			par10{0}, ans10{1},
		},
	}
	fmt.Println("______offer10 青蛙跳台阶 简单题")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了22.93%的用户
	递归改成迭代，效率会更高
	`)
	for _, example := range qs {
		par, as := example.par10, example.ans10

		res := numWays(par.n)
		// res_dfs :=numWays_DFS(par.n,par.relation,par.k)
		if res != as.output {
			t.Errorf("error res=%v,ans=%v\n", res, as.output)
		} else {
			fmt.Printf("pass \n")
		}
	}

}
