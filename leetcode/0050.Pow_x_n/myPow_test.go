package leetcode

import (
	"fmt"
	"math"
	"testing"
)
type question50 struct{
	par50
	ans50
}
type par50 struct{
	x float64
	n int
}
type ans50 struct{
	output float64
}

func Test_MyPow(t *testing.T){
	qs:=[]question50{
		{
			par50{2.00000,10},
			ans50{1024.00000},
		},
		{
			par50{2.10000,3},
			ans50{9.26100},
		},
		{
			par50{2.00000,-2},
			ans50{0.25000},
		},
		{
			par50{0.00001,2147483647},
			ans50{0},
		},
		{
			par50{
				-1.00000,
				2147483647},
			ans50{-1},
		},

	}
	fmt.Println("______________50 Pow(x,n) 中等题")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了99.88%的用户`)
	for _,ex := range qs{
		par,as := ex.par50,ex.ans50
		res:=myPow(par.x,par.n)
		if !isEqual_Float64(res,as.output){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Println("pass")
		}
	}
}

func isEqual_Float64(a,b float64)bool{
	return math.Abs(a-b)<1e-7
}