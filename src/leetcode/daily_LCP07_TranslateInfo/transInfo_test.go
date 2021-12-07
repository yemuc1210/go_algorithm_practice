package leetcode

import (
	"fmt"
	"testing"
)

type questionLCP07 struct{
	par07
	ans07
}
type par07 struct{
	 relation [][]int
	 n int
	 k int
}
type ans07 struct{
	nums int
}


func Test_TransInfo(t *testing.T){
	qs:=[]questionLCP07{
		{
			par07{n:5,k: 3,relation: [][]int{
				{0,2},{2,1},{3,4},{2,3},{1,4},{2,0},{0,4},
			}},
			ans07{3},
		},
		{
			par07{n: 3,k: 2,relation: [][]int{
				{0,2},{2,1},
			}},
			ans07{0},
		},
	}
	fmt.Println("______LCP 07 传递信息 简单题")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了100%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs{
		par,as := example.par07,example.ans07

		res:=numWays(par.n,par.relation,par.k)
		res_dfs :=numWays_DFS(par.n,par.relation,par.k)
		if res != as.nums && res_dfs!=as.nums{
			t.Errorf("error res=%v,ans=%v\n",res,as.nums)
		}else{
			fmt.Printf("pass \n")
		}
	}

}