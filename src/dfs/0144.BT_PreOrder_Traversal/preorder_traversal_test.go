package leetcode

import (
	"fmt"
	"structs"
	"testing"
)
type question144 struct {
	para144
	ans144
}

// para 是参数
// one 代表第一个参数
type para144 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans144 struct {
	one []int
}

func Test_Problem94(t *testing.T) {

	qs := []question144{

		{
			para144{[]int{}},
			ans144{[]int{}},
		},

		{
			para144{[]int{1}},
			ans144{[]int{1}},
		},

		{
			para144{[]int{1, structs.NULL, 2, 3}},
			ans144{[]int{1, 2, 3}},
		},
	}

	fmt.Println("____________测试144 二叉树先序遍历（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了66.96%的用户`)
	for _, q := range qs {
		_, p := q.ans144, q.para144
		fmt.Printf("【input】:%v      ", p)
		root := structs.Ints2TreeNode(p.one)
		// res := inorderTraversal_1(root)
		res1 := inorderTraversal(root)
		if  !isEqual_arr(res1,q.ans144.one){
			t.Errorf("error\n")
		}else{
			fmt.Printf("【output】:%v      \n", res1)
		}
	}
	fmt.Printf("\n\n\n")
}

func isEqual_arr(a []int,b []int)bool{
	if len(a) != len(b){
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i] != b[i]{
			return false
		}
	}
	return true
}
