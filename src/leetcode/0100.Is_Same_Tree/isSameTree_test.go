package leetcode

import (
	"fmt"
	"testing"
	"structs"
)

type question100 struct {
	para100
	ans100
}

// para 是参数
// one 代表第一个参数
type para100 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans100 struct {
	one bool
}

func Test_Problem100(t *testing.T) {

	qs := []question100{

		{
			para100{[]int{}, []int{}},
			ans100{true},
		},

		{
			para100{[]int{}, []int{1}},
			ans100{false},
		},

		{
			para100{[]int{1}, []int{1}},
			ans100{true},
		},

		{
			para100{[]int{1, 2, 3}, []int{1, 2, 3}},
			ans100{true},
		},

		{
			para100{[]int{1, 2}, []int{1, structs.NULL, 2}},
			ans100{false},
		},

		{
			para100{[]int{1, 2, 1}, []int{1, 1, 2}},
			ans100{false},
		},
	}
	fmt.Println("____________测试100 二叉树是否相同（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了86.71%的用户`)
	for _, q := range qs {
		as, par := q.ans100, q.para100
		fmt.Printf("【input】:%v      ", par)
		rootOne := structs.Ints2TreeNode(par.one)
		rootTwo := structs.Ints2TreeNode(par.two)
		res := isSameTree(rootOne,rootTwo)
		if res!=as.one{
			t.Errorf("error\n")
		}else{
			fmt.Printf("pass 【output】:%v      \n", res)
		}
	}
	fmt.Printf("\n\n\n")
}