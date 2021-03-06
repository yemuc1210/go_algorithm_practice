package offer38

import (
	"fmt"
	"testing"
)

type question46 struct {
	para46
	ans46
}

// para 是参数
// one 代表第一个参数
type para46 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans46 struct {
	one []string
}

func Test_Problem46(t *testing.T) {

	qs := []question46{

		{
			para46{"abc"},
			ans46{[]string{"abc","acb","bac","bca","cab","cba"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 46------------------------\n")

	for _, q := range qs {
		_, p := q.ans46, q.para46
		fmt.Printf("【input】:%v       【output】:%v\n", p, permutation(p.s))
	}
	fmt.Printf("\n\n\n")
}
