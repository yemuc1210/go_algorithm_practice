package leetcode

import (
	"fmt"
	"testing"
)
type question51 struct {
	para51
	ans51
}

// para 是参数
// one 代表第一个参数
type para51 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans51 struct {
	one [][]string
}

func Test_Problem51(t *testing.T) {

	qs := []question51{

		{
			para51{4},
			ans51{[][]string{
				{".Q..",
					"...Q",
					"Q...",
					"..Q."},
				{"..Q.",
					"Q...",
					"...Q",
					".Q.."},
			}},
		},
	}

	fmt.Println("______________51 N皇后问题 困难")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了95.78%的用户
	内存消耗：3.1 MB, 在所有 Go 提交中击败了97.89%的用户`)
	for _,ex := range qs{
		par,as := ex.para51,ex.ans51
		res:=solveNQueens(par.one)
		if !isEqual_str(res,as.one){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.one)
		}else{
			fmt.Println("pass")
		}
	}
}
func isEqual_str(a,b [][]string)bool{
	if len(a)!=len(b){
		return false
	}
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[i]);j++{
			if a[i][j]!=b[i][j]{
				return false
			}
		}
	}
	return true
}