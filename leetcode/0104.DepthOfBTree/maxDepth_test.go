package lt104

import (
	"fmt"
	"go_practice/structs"
	"testing"
)
type question104 struct{
	para104
	ans104
}

// para 是参数
// one 代表第一个参数
type para104 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans104 struct {
	one int
}
func Test_Problem104(t *testing.T) {

	qs := []question104{

		{
			para104{[]int{}},
			ans104{0},
		},

		{
			para104{[]int{3, 9, 20, structs.NULL, structs.NULL, 15, 7}},
			ans104{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 104------------------------\n")

	for _, q := range qs {
		as, p := q.ans104, q.para104
		fmt.Printf("【input】:%v      ", p)
		root := structs.Ints2TreeNode(p.one)
		res := maxDepth(root)
		res1 := maxDepth1(root)
		if res!=as.one || res1!=as.one{
			t.Errorf("error\n")
		}else{
			fmt.Printf("pass 【output】:%v      \n", res)
		}
	}

}

