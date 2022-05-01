package offer13

import (
	"fmt"
	"testing"
)
type question13 struct{
	par13
	ans13
}
type par13 struct{
	m,n,k int
}
type ans13 struct{
	out int
}
func Test_Pro13(t *testing.T){
	qs:=[]question13{
		{
			par13{1,2,3},
			ans13{2},
		},
		{
			par13{2,3,1},
			ans13{3},
		},
		{
			par13{3,1,0},
			ans13{1},
		},
		{
			par13{38,15,9},
			ans13{135},
		},
	}
	fmt.Println("_________________13 机器人的运动范围 中等题")
	fmt.Println(`
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了88.13%的用户
	`)
	for _,ex := range qs{
		par,as := ex.par13,ex.ans13
		res:=movingCount(par.m,par.n,par.k)
		if res!=as.out{
			t.Errorf("error par=%v,res=%v,ans=%v\n",par,res,as.out)
		}else{
			fmt.Printf("pass \n")
		}
	}
}