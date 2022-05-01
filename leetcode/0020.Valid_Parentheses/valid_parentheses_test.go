package leetcode

import (
	"fmt"
	"testing"
)

type question20 struct{
	par20
	ans20
}

type par20 struct{
	input string
}

type ans20 struct{
	output bool
}

func Test_Valid_Parenthese(t *testing.T){
	qs := []question20{
		{
			par20{"()"},
			ans20{true},
		},
		{
			par20{"(){}[]"},
			ans20{true},
		},
		{
			par20{"{(])"},
			ans20{false},
		},
		{
			par20{"(]"},
			ans20{false},
		},
		{
			par20{"([)]"},
			ans20{false},
		},
	}
	fmt.Println("____________测试20 有效括号（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了26.09%的用户`)
	for _,example := range qs {
		par,as := example.par20,example.ans20

		res := isValid(par.input)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par.input,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par.input,res)
		}
	}

}