package leetcode

import (
	"fmt"
	"testing"
)
type question52 struct {
	para52
	ans52
}

// para 是参数
// one 代表第一个参数
type para52 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans52 struct {
	one int
}

func Test_Problem51(t *testing.T) {

	qs := []question52{

		{
			para52{1},
			ans52{1},
		},

		{
			para52{2},
			ans52{0},
		},
		{
			para52{3},
			ans52{0},
		},

		{
			para52{4},
			ans52{2},
		},

		{
			para52{5},
			ans52{10},
		},

		{
			para52{6},
			ans52{4},
		},

		{
			para52{7},
			ans52{40},
		},

		{
			para52{8},
			ans52{92},
		},

		{
			para52{9},
			ans52{352},
		},

		{
			para52{10},
			ans52{724},
		},
	}

	fmt.Println("______________52 N皇后问题 困难")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了31.22%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了16.03%的用户`)
	for _,ex := range qs{
		par,as := ex.para52,ex.ans52
		res:=totalNQueens(par.one)
		res1:=totalNQueens1(par.one)
		if res!=as.one || res1!=as.one{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.one)
		}else{
			fmt.Println("pass")
		}
	}
}
