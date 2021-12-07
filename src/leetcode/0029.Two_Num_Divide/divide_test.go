package leetcode

import (
	"fmt"
	"testing"
)

type question29 struct {
	para29
	ans29
}

// para 是参数
// one 代表第一个参数
type para29 struct {
	dividend int
	divisor  int
}

// ans 是答案
// one 代表第一个答案
type ans29 struct {
	one int
}

func Test_Problem29(t *testing.T) {

	qs := []question29{

		{
			para29{10, 3},
			ans29{3},
		},

		{
			para29{7, -3},
			ans29{-2},
		},

		{
			para29{-1, 1},
			ans29{-1},
		},

		{
			para29{1, -1},
			ans29{-1},
		},

		{
			para29{2147483647, 3},
			ans29{715827882},
		},
	}

	fmt.Printf("------------------------测试 29 除法 中等------------------------\n")
	fmt.Println("____________测试 29 除法（中等题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2.4 MB, 在所有 Go 提交中击败了64.95%的用户`)
	for _,example := range qs{
		par,as := example.para29,example.ans29
		res:=divide(par.dividend,par.divisor)
		if res != as.one{
			t.Errorf("error [input]:%v   [output]:%v\n",par,res)
		}else{
			fmt.Printf("pass [input]:%v   [output]:%v\n",par,res)
		}
	}
}
