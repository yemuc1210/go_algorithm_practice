package leetcode

import (
	"fmt"
	"go_practice/structs"
	"testing"
)

type question145 struct {
	para145
	ans145
}

// para 是参数
// one 代表第一个参数
type para145 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans145 struct {
	one []int
}

func Test_Problem94(t *testing.T) {

	qs := []question145{

		{
			para145{[]int{}},
			ans145{[]int{}},
		},

		{
			para145{[]int{1}},
			ans145{[]int{1}},
		},

		{
			para145{[]int{1, structs.NULL, 2, 3}},
			ans145{[]int{3, 2, 1}},
		},
	}

	fmt.Println("____________测试145 二叉树后序遍历（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了80.25%的用户`)
	for _, q := range qs {
		_, p := q.ans145, q.para145
		fmt.Printf("【input】:%v      ", p)
		root := structs.Ints2TreeNode(p.one)
		res := postorderTraversal(root)

		if !isEqual_arr(res, q.ans145.one) {
			t.Errorf("error out:%v\n",res)
		} else {
			fmt.Printf("【output】:%v      \n", res)
		}
	}
}

func isEqual_arr(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
