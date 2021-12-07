package offer04

import (
	"fmt"
	"testing"
)
type questioon04 struct{
	par04
	ans04
}
type par04 struct{
	matrix [][]int
	target int
}
type ans04 struct{
	result bool
}

func Test_Find2d(t *testing.T){
	qs:=[]questioon04{
		{
			par04{[][]int{
				{1,   4,  7, 11, 15},
  {2,   5,  8, 12, 19},
  {3,   6,  9, 16, 22},
  {10, 13, 14, 17, 24},
  {18, 21, 23, 26, 30},
			},5},
			ans04{true},
		},
		{
			par04{[][]int{
				{5},{6},
			},6},
			ans04{true},
		},
	}
	fmt.Println("_________________04 二维数组的查找 简单题")
	fmt.Println(`
	执行用时：28 ms, 在所有 Go 提交中击败了90.43%的用户
	内存消耗：6.7 MB, 在所有 Go 提交中击败了91.76%的用户
	`)
	for _,example := range qs{
		par,as := example.par04,example.ans04

		res := findNumberIn2DArray(par.matrix,par.target)
		if res!=as.result{
			t.Errorf("error\n")
		}else{
			fmt.Println("pass")
		}
	}
}