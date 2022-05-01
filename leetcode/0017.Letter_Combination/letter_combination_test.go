package leetcode

import (
	"fmt"
	"testing"
)

type question17 struct{
	par17
	ans17
}

type par17 struct{
	input string
}
type ans17 struct{
	output []string
}

func Test_Letter_Combination(t *testing.T){
	qs:=[]question17{
		{
			par17{"23"},
			ans17{[]string{
				"ad","ae","af","bd","be","bf","cd","ce","cf",
			}},
		},
		{
			par17{""},
			ans17{[]string{}},
		},
		{
			par17{"2"},
			ans17{[]string{"a","b","c"}},
		},
	}
	fmt.Println("____________测试17 字母组合（中等题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了99.57%的用户`)
	for _,example := range qs{
		par,as := example.par17,example.ans17
		res:=letterCombinations(par.input)
		if !isEqual_arr(res,as.output){
			t.Errorf("error [input]:%v   [output]:%v\n",par.input,res)
		}else{
			fmt.Printf("pass [input]:%v   [output]:%v\n",par.input,res)
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