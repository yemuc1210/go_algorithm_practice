package fib

import (
	"fmt"
	"testing"
)


type question10 struct{
	par10
	ans10
}
type par10 struct{
	n int
}
type ans10 struct{
	fib_n int
}
func Test_Fib(t *testing.T){
	qs:=[]question10{
	{
		par10{2},ans10{1},
	},
	{
		par10{5},ans10{5},
	},
	}
	fmt.Println("__________________0010 fib 简单题")
	fmt.Println(`
	执行用时 0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：1.9 MB, 在所有 Go 提交中击败了75.31%的用户
	`)
	for _,example:=range qs{
		par,as := example.par10,example.ans10
		res:=fib(par.n)
		res1:=fib_1(par.n)
		if res !=as.fib_n && res1 != as.fib_n{
			t.Errorf("error\n")
		}else{
			fmt.Printf("pass\n")
		}
	}

}