package lt34

import (
	"fmt"
	"testing"
)

type question34 struct {
	para34
	ans34
}

// para 是参数
// one 代表第一个参数
type para34 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans34 struct {
	one []int
}

func Test_Problem34(t *testing.T) {

	qs := []question34{

		{
			para34{[]int{5, 7, 7, 8, 8, 10}, 8},
			ans34{[]int{3, 4}},
		},

		{
			para34{[]int{5, 7, 7, 8, 8, 10}, 6},
			ans34{[]int{-1, -1}},
		},
		{
			para34{[]int{1},1},
			ans34{[]int{0,0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 34------------------------\n")
	fmt.Println(`执行结果：
	通过
	执行用时：8 ms, 在所有 Go 提交中击败了87.64%的用户
	内存消耗：3.9 MB, 在所有 Go 提交中击败了60.33%的用户`)
	for _, example := range qs {
		as,par := example.ans34, example.para34

		res := searchRange(par.nums,par.target)
		if !isEqual(res, as.one) {
			t.Errorf("error [input]:%v   [output]:%v\n", par, res)
		} else {
			fmt.Printf("pass [input]:%v   [output]:%v\n", par, res)
		}
		// fmt.Printf("【input】:%v       【output】:%v\n", p, searchRange(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
func isEqual(a []int, b []int) bool {
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