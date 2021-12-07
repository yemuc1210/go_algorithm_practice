package lt931

import (
	"fmt"
	"testing"
)

type question931 struct{
	matrix [][]int
	ans int
}

func Test_Problem931(t *testing.T){
	qs:=[]question931{
		{
			matrix: [][]int{
				{2,1,3},{6,5,4},{7,8,9},
			},
			ans: 13,
		},
		{
			matrix: [][]int{
				{-19,57},{-40,-5},
			},
			ans: -59,
		},
	}
	for _,ex := range qs{
		par,as := ex.matrix,ex.ans
		res := minFallingPathSum(par)
		if res!=as{
			t.Errorf("error out=%v  ans=%v\n",res,as)
		}else{
			fmt.Println("pass")
		}
	}
}