package offer12

import (
	"fmt"
	"testing"
)

type question12 struct{
	par12
	ans12
}
type par12 struct{
	board [][]byte
	word string
}
type ans12 struct{
	result bool
}

func Test_Problem12(t *testing.T){
	qs:=[]question12{
		{
			par12{board: [][]byte{
				{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'},
			},word: "ABCCED"},
			ans12{true},
		},
		{
			par12{board: [][]byte{
				{'a','b'},{'c','d'},
			},word: "abcd"},
			ans12{false},
		},
	}
	fmt.Println("_________________12 矩阵中的路径 中等题")
	fmt.Println(`
	执行用时：4 ms, 在所有 Go 提交中击败了95.57%的用户
	内存消耗：3.3 MB, 在所有 Go 提交中击败了54.25%的用户
	`)
	for _,example := range qs{
		par,as := example.par12,example.ans12

		res := exist(par.board,par.word)
		if res!=as.result{
			t.Errorf("error\n")
		}else{
			fmt.Println("pass")
		}
	}
}