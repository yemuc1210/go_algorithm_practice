package lt46

import (
	"fmt"
	"testing"
)

type question46 struct {
	para46
	ans46
}

// para 是参数
// one 代表第一个参数
type para46 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans46 struct {
	one [][]int
}

func Test_Problem46(t *testing.T) {

	qs := []question46{

		{
			para46{[]int{1, 2, 3}},
			ans46{[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		},
	}

	fmt.Println("______________52 N皇后问题 困难")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2.7 MB, 在所有 Go 提交中击败了33.20%的用户`)
	for _,ex := range qs{
		// par,_ := ex.para46,ex.ans46
		// res:=permute(par.s)


		fmt.Printf("pass %v    %v",ex.ans46.one,permute(ex.para46.s) )
		
	}
}
