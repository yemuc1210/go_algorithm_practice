package leetcode

import (
	"fmt"
	"testing"
)

type question773 struct{
	par773
	ans773
}
type par773 struct{
	input [][]int
}
type ans773 struct{
	output int
}

func Test_Sliding_Puzzle(t *testing.T){
	qs:=[]question773{
		{
			par773{[][]int{
				{1,2,3},{4,0,5},
			}},
			ans773{1},
		},
		{
			par773{[][]int{
				{1,2,3},{5,4,0},
			}},
			ans773{-1},
		},
		{
			par773{[][]int{
				{4,1,2},{5,0,3},
			}},
			ans773{5},
		},
		{
			par773{[][]int{
				{3,2,4},{1,5,0},}},
			ans773{14},
		},
	}

	fmt.Println("____________测试773 滑动谜题（困难题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了81.82%的用户
	内存消耗：3.6 MB, 在所有 Go 提交中击败了45.45%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs {
		par,as := example.par773,example.ans773

		res := slidingPuzzle_1(par.input)
		if res!=as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}