package leetcode

import (
	"fmt"
	"testing"
)

type question7 struct {
	para7
	ans7
}

// para 是参数
// one 代表第一个参数
type para7 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans7 struct {
	one int
}

func Test_Problem7(t *testing.T) {

	qs := []question7{

		{
			para7{321},
			ans7{123},
		},

		{
			para7{-123},
			ans7{-321},
		},

		{
			para7{120},
			ans7{21},
		},

		{
			para7{1534236469},
			ans7{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 7------------------------\n")

	
	for _, q := range qs {
		as, p := q.ans7, q.para7

		ans := reverse7(p.one)
		if as.one != ans{
			t.Errorf("error [input]:%v      [output]:%v\n",p.one,ans)
		}

		fmt.Printf("【input】:%v    【output】:%v\n", p.one, ans)
	}
	fmt.Printf("\n\n\n")
}
