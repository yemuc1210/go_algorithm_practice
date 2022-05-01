package leetcode

import (
	"fmt"
	"testing"
)

type question909 struct{
	par909
	ans909
}

type par909 struct{
	input [][]int
}
type ans909 struct{
	output int
}

func Test_Snake_Ladders(t *testing.T){
	qs:=[]question909{
		{
			par909{[][]int{
				{-1,-1,-1,-1,-1,-1},
				{-1,-1,-1,-1,-1,-1},
				{-1,-1,-1,-1,-1,-1},
				{-1,35,-1,-1,13,-1},
				{-1,-1,-1,-1,-1,-1},
				{-1,15,-1,-1,-1,-1},
			}},
			ans909{4},
		},
	}
	fmt.Println("____________测试909 蛇梯（中等题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：16 ms, 在所有 Go 提交中击败了78.57%的用户
	内存消耗：5.1 MB, 在所有 Go 提交中击败了92.86%的用户`)
	for _,example := range qs{
		par,as := example.par909,example.ans909
		res:=snakesAndLadders(par.input)
		if res!=as.output{
			t.Errorf("error [input]:%v   [output]:%v\n",par.input,res)
		}else{
			fmt.Printf("pass [input]:%v   [output]:%v\n",par.input,res)
		}
	}
}