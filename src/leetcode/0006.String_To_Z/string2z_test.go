package leetcode

import (
	"fmt"
	"testing"
)

type question6 struct{
	para6
	ans6
}

type para6 struct{
	s string
	numRows int
}

type ans6 struct{
	output string
}

func Test_string2z(t *testing.T){
	qs := []question6{
		{
			para6{
				s:"PAYPALISHIRING",
				numRows: 3,
			},
			ans6{"PAHNAPLSIIGYIR"},
		},
		{
			para6{s : "PAYPALISHIRING", numRows:3},
			ans6{"PAHNAPLSIIGYIR"},
		},
		{
			para6{s : "PAYPALISHIRING", numRows : 4},
			ans6{"PINALSIGYAHRPI"},
		},
		{
			para6{s : "A", numRows :1},
			ans6{"A"},
		},
		{
			para6{s : "AB", numRows :1},
			ans6{"AB"},
		},
	}

	fmt.Println("____________测试6 z型____________")
	for _,ques := range qs {
		par,as := ques.para6,ques.ans6

		res := convert(par.s,par.numRows)
		if res != as.output {
			t.Errorf("error [input]:%v     [output]:%v\n",par,res)
		}else{
			fmt.Printf("pass [input]:%v    [output]:%v\n",par,res)
		}

	}
	fmt.Println(`
		执行用时：628 ms, 在所有 Go 提交中击败了5.46%的用户
		内存消耗：74.1 MB, 在所有 Go 提交中击败了5.02%的用户`)
}