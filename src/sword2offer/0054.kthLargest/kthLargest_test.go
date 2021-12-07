package offer54

import (
	"fmt"
	"structs"
	"testing"
)
type question54 struct {
	para54
	ans54
}

// para 是参数
// one 代表第一个参数
type para54 struct {
	one []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans54 struct {
	one int
}

func Test_Problem94(t *testing.T) {

	qs := []question54{

		// {
		// 	para54{[]int{},0},
		// 	ans54{},
		// },

		{
			para54{[]int{5,3,6,2,4, structs.NULL,structs.NULL, 1},3},
			ans54{4},
		},

		{
			para54{[]int{3,1,4, structs.NULL, 2},1},
			ans54{4},
		},
	}

	fmt.Println("____________测试894 二叉树中序遍历（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了66.96%的用户`)
	for _, q := range qs {
		as, p := q.ans54, q.para54
		fmt.Printf("【input】:%v      ", p)
		root := structs.Ints2TreeNode(p.one)
		res := kthLargest(root,p.k)
		if res!=as.one{
			t.Errorf("error\n")
		}else{
			fmt.Printf("pass 【output】:%v      \n", res)
		}
	}
}
