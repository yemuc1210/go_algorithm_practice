package offer55

import (
	"fmt"
	"structs"
	"testing"
)
type question55 struct{
	para55
	ans55
}

// para 是参数
// one 代表第一个参数
type para55 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans55 struct {
	one int
}
func Test_Problem104(t *testing.T) {

	qs := []question55{

		{
			para55{[]int{}},
			ans55{0},
		},

		{
			para55{[]int{3, 9, 20, structs.NULL, structs.NULL, 15, 7}},
			ans55{3},
		},
	}

	// fmt.Printf("------------------------------------------------\n")

	for _, q := range qs {
		as, p := q.ans55, q.para55
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

