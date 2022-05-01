package leetcode

import (
	"fmt"
	"testing"
)

type question26 struct{
	par26
	ans26
}

type par26 struct{
	input []int
}
type ans26 struct{
	output int
}

func Test_Remove_Duplicates(t *testing.T){
	qs:=[]question26{
		{
			par26{[]int{1,1}},
			ans26{1},
		},
		{
			par26{[]int{1,1,2}},
			ans26{2},
		},
		{
			par26{[]int{0,0,1,1,2,2,3,3,4}},
			ans26{5},
		},
	}
	fmt.Println("____________测试26 删除有序数组中的重复项（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：12 ms, 在所有 Go 提交中击败了31.25%的用户
	内存消耗：4.3 MB, 在所有 Go 提交中击败了100%的用户`)
	for _,example := range qs {
		par,as := example.par26,example.ans26

		res := removeDuplicates(par.input)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par.input,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par.input,res)
		}
	}
}