package leetcode

import (
	"fmt"
	"testing"
)

type question8 struct{
	para8
	ans8
}

type para8 struct{
	input string
}

type ans8 struct{
	output int
}

func Test_String2Int(t *testing.T){
	qs:= []question8{
		{
			para8{"42"},
			ans8{42},
		},
		{
			para8{"   -42"},
			ans8{-42},
		},
		{
			para8{"4193 with words"},
			ans8{4193},
		},
		{
			para8{"words and 987"},
			ans8{0},
		},
		{
			para8{"-91283472332"},
			ans8{-2147483648},
		},
	}
	
	fmt.Println("__________测试8  字符串转数字 ______________")
	for _,example := range qs{
		ans,par := example.ans8,example.para8

		res := my_atoi(par.input)

		if res != ans.output {
			t.Errorf("error [input]:%v       [output]:%v\n",par.input,res)
		}else{
			fmt.Printf("pass [input]:%v       [output]:%v\n",par.input,res)
		}
	}
	fmt.Println(`
	执行用时：4 ms, 在所有 Go 提交中击败了48.48%的用户
	内存消耗：2.2 MB, 在所有 Go 提交中击败了70.28%的用户`)
}