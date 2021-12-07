package leetcode

import (
	"fmt"
	"testing"
)


type question27 struct{
	par27
	ans27
}

type par27 struct{
	nums []int
	val int
}
type ans27 struct{
	output int
}

func Test_Remove_Element(t *testing.T){
	qs :=[]question27{
		{
			par27{nums:[]int{1,1},val: 1},
			ans27{0},
		},
		{
			par27{nums:[]int{1,1,2},val:1},
			ans27{1},
		},
		{
			par27{nums:[]int{0,0,1,1,2,2,3,3,4},val: 3},
			ans27{7},
		},
		{
			par27{nums:[]int{3,2,2,3},val:3},
			ans27{2},
		},
		{
			par27{nums:[]int{0,1,2,2,3,0,4,2},val: 2},
			ans27{5},
		},
	}
	fmt.Println("____________测试27 删除元素（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了59.49%的用户`)
	for _,example := range qs {
		par,as := example.par27,example.ans27

		res := removeElement(par.nums,par.val)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}