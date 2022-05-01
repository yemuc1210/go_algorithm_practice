package lt110

import (
	"fmt"
	"go_practice/structs"
	"testing"
)

type question110 struct{
	par []int
	ans bool
}
func Test_isBalanced(t *testing.T){
	qs:=[]question110{
		{
			par: []int{3, 4, 4, 5, structs.NULL, structs.NULL, 5, 6, structs.NULL, structs.NULL, 6},
			ans: false,
		},
		{
			par: []int{1,2,2,structs.NULL,3,3},
			ans: true,
		},
		{
			par: []int{},
			ans: true,
		},
		{
			par: []int{1},
			ans: true,
		},
		{
			par: []int{1,2,3},
			ans: true,
		},
		{
			par: []int{1, 2, 2, 3, 4, 4, 3},
			ans: true,
		},
	}
	for _,ex := range qs{
		par,as := ex.par,ex.ans
		root := structs.Ints2TreeNode(par)
		res := isBalanced(root)
		if res!= as{
			t.Errorf("error [input]:%v [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}