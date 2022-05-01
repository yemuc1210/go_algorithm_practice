package leetcode

import (
	"fmt"
	"testing"
)

type question35 struct{
	par35
	ans35
}
type par35 struct{
	nums []int
	target int 
}
type ans35 struct{
	output int
}

func Test_Search_Insert(t *testing.T){
	qs :=[]question35{
		{
			par35{[]int{1,3,5,6},5},
			ans35{2},
		},
		{
			par35{[]int{1,3,5,6},2},
			ans35{1},
		},
		{
			par35{[]int{1,3,5,6},7},
			ans35{4},
		},
		{
			par35{[]int{1,3,5,6},0},
			ans35{0},
		},
	}

	
	fmt.Println("____________测试27 删除元素（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了91.14%的用户
	内存消耗：3 MB, 在所有 Go 提交中击败了56.72%的用户`)
	for _,example := range qs {
		par,as := example.par35,example.ans35

		res := searchInsert(par.nums,par.target)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}