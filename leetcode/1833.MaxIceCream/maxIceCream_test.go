package leetcode

import (
	"fmt"
	"testing"
)
type question1833 struct{
	par1833
	ans1833
}
type par1833 struct{
	costs []int
	coins int
}
type ans1833 struct{
	nums int
}

func Test_MaxIceCream(t *testing.T){
	qs :=[]question1833{
		{
			par1833{[]int{1,3,2,4,1},7},
			ans1833{4},
		},
		{
			par1833{[]int{10,6,8,7,7,8},5},
			ans1833{0},
		},
		{
			par1833{[]int{1,6,3,1,2,5},20},
			ans1833{6},
		},
	}
	fmt.Println("______1833 雪糕的最大数量 中等题")
	fmt.Println(`执行结果：
	通过
	执行用时：224 ms, 在所有 Go 提交中击败了63.96%的用户
	内存消耗：9.2 MB, 在所有 Go 提交中击败了69.07%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs{
		par,as := example.par1833,example.ans1833

		res:=maxIceCream(par.costs,par.coins)
		// res_dfs :=numWays_DFS(par.n,par.relation,par.k)
		if res != as.nums {
			t.Errorf("error res=%v,ans=%v\n",res,as.nums)
		}else{
			fmt.Printf("pass \n")
		}
	}
}