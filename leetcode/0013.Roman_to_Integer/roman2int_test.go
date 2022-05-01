package leetcode

import (
	"fmt"
	"testing"
)

type question13 struct {
	para13
	ans13
}

// para 是参数
// one 代表第一个参数
type para13 struct {
	input string
}

// ans 是答案
// one 代表第一个答案
type ans13 struct {
	output int
}



func Test_Roman2Int(t *testing.T){
	qs := []question13{

		{
			para13{"III"},
			ans13{3},
		},

		{
			para13{"IV"},
			ans13{4},
		},

		{
			para13{"IX"},
			ans13{9},
		},

		{
			para13{"LVIII"},
			ans13{58},
		},

		{
			para13{"MCMXCIV"},
			ans13{1994},
		},

		{
			para13{"MCMXICIVI"},
			ans13{2014},
		},
	}
	fmt.Println("_______________测试13  罗马字母转数字 从左往右___________")

	for _,ques := range qs {
		par,ans := ques.para13,ques.ans13

		res := roman2int(par.input)

		if res != ans.output {
			t.Errorf("error [input]:%v       [output]:%v\n",par.input,res)
		}else{
			fmt.Printf(" pass [input]:%v    [output]:%v\n",par.input,res)
		}
	}
}

func Test_RomanToInt(t *testing.T){
	qs := []question13{

		{
			para13{"III"},
			ans13{3},
		},

		{
			para13{"IV"},
			ans13{4},
		},

		{
			para13{"IX"},
			ans13{9},
		},

		{
			para13{"LVIII"},
			ans13{58},
		},

		{
			para13{"MCMXCIV"},
			ans13{1994},
		},

		{
			para13{"MCMXICIVI"},
			ans13{2014},
		},
	}
	fmt.Println("_______________测试13  罗马字母转数字 从右往左___________")

	for _,ques := range qs {
		par,ans := ques.para13,ques.ans13

		res := romanToInt(par.input)

		if res != ans.output {
			t.Errorf("error [input]:%v       [output]:%v\n",par.input,res)
		}else{
			fmt.Printf(" pass [input]:%v    [output]:%v\n",par.input,res)
		}
	}
}