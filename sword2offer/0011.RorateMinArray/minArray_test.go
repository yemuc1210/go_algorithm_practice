package offer11

import (
	"fmt"
	"testing"
)
type question11 struct{
	para11
	ans11
}

type para11 struct{
	nums []int
}
type ans11 struct{
	one int
}

func Test_Problem11(t *testing.T){
	qs := []question11{

		{
			para11{[]int{10, 2, 10, 10, 10, 10}},
			ans11{2},
		},

		{
			para11{[]int{1, 1}},
			ans11{1},
		},

		{
			para11{[]int{2, 2, 2, 0, 1}},
			ans11{0},
		},

		{
			para11{[]int{5, 1, 2, 3, 4}},
			ans11{1},
		},

		{
			para11{[]int{1}},
			ans11{1},
		},

		{
			para11{[]int{1, 2}},
			ans11{1},
		},

		{
			para11{[]int{2, 1}},
			ans11{1},
		},

		{
			para11{[]int{2, 3, 1}},
			ans11{1},
		},

		{
			para11{[]int{1, 2, 3}},
			ans11{1},
		},

		{
			para11{[]int{3, 4, 5, 1, 2}},
			ans11{1},
		},

		{
			para11{[]int{4, 5, 6, 7, 0, 1, 2}},
			ans11{0},
		},
	}

	fmt.Println("_________________11 旋转数组中的最小元素 简单题")
	fmt.Println(`
	执行用时：4 ms, 在所有 Go 提交中击败了91.13%的用户
	内存消耗：3.1 MB, 在所有 Go 提交中击败了56.82%的用户
	`)
	for _,example := range qs{
		par,as := example.para11,example.ans11

		res := minArray(par.nums)
		res1:=minArray1(par.nums)
		if res!=as.one || res1!=as.one{
			t.Errorf("error\n")
		}else{
			fmt.Println("pass")
		}
	}
}