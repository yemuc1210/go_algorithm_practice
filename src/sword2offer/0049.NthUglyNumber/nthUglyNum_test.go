package offer49

import (
	"fmt"
	"testing"
)

type question49 struct {
	para49
	ans49
}

// para 是参数
// one 代表第一个参数
type para49 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans49 struct {
	one int
}

func Test_Problem49(t *testing.T) {

	qs := []question49{
		{
			para49{10},
			ans49{12},
		},
		
	}


	for _, ex := range qs {
		par,as := ex.para49,ex.ans49
		res := nthUglyNumber(par.one)
		if res!=as.one{
			t.Errorf("error [input]:%v  [out]:%v   [ans]:%v\n",par.one,res,as.one)
		}else{
			fmt.Println("pass")
		}
	}
	fmt.Printf("\n\n\n")
}
