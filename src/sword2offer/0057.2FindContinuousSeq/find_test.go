package offer57_2

import (
	"fmt"

	"testing"
)

type question57_2 struct{
	target int
	res [][]int
}

func Test_Find(t *testing.T){
	qs :=[]question57_2{
		{
			target: 9,
			res: [][]int{{2,3,4},{4,5}},
		},
		{
			target: 15,
			res: [][]int{{1,2,3,4,5},{4,5,6},{7,8}},
		},
	}
	for _,ex := range qs{
		par,as := ex.target,ex.res
		res := findContinuousSequence(par)
		if !isEqual_arr(res,as){
			t.Errorf("error [input]:%v  [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}
func isEqual_arr(a,b [][]int)bool{
	if len(a)!=len(b) {
		return false
	}
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a[i]);j++{
			if a[i][j]!=b[i][j]{
				return false
		
			}
		}
	}
	return true
}