package leetcode

import (
	"fmt"
	"testing"
)

type question22 struct{
	par22
	ans22
}

type par22 struct{
	input int
}

type ans22 struct{
	output []string
}

func Test_Valid_Parenthese(t *testing.T){
	qs := []question22{
		{
			par22{3},
			ans22{[]string{"((()))","(()())","(())()","()(())","()()()",}},
		},
		{
			par22{1},
			ans22{[]string{"()"}},
		},
		
	}
	fmt.Println("____________测试20 有效括号（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2.7 MB, 在所有 Go 提交中击败了84.67%的用户`)
	for _,example := range qs {
		par,as := example.par22,example.ans22

		res := generateParenthesis(par.input)
		if !isEqual_arr(res,as.output){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par.input,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par.input,res)
		}
	}

}

func isEqual_arr(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}