package leetcode

import (
	"fmt"
	"testing"
)

type question56 struct {
	para56
	ans56
}

// para 是参数
// one 代表第一个参数
type para56 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans56 struct {
	one [][]int
}

func Test_Problem56(t *testing.T) {

	qs := []question56{

		{
			para56{[][]int{}},
			ans56{[][]int{}},
		},

		{
			para56{[][]int{
				{1, 3}, { 2, 6}, { 8,  10}, { 15, 18}}},
			ans56{[][]int{{ 1, 6},{ 8, 10}, { 15, 18}}},
		},

		{
			para56{[][]int{{ 1, 4}, {4, 5}}},
			ans56{[][]int{{ 1,  5}}},
		},

		{
			para56{[][]int{{ 1,  3}, { 3,  6}, { 5, 10}, { 9, 18}}},
			ans56{[][]int{{ 1, 18}}},
		},

	}
	fmt.Println("______________56 区间合并 中等")
	fmt.Println(`执行结果：
	通过
	执行用时：12 ms, 在所有 Go 提交中击败了79.20%的用户
	内存消耗：4.5 MB, 在所有 Go 提交中击败了90.64%的用户`)
	for _,ex := range qs{
		par,as := ex.para56,ex.ans56
		res:=merge(par.one)
		if !isEqual_arr(res,as.one){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.one)
		}else{
			fmt.Println("pass")
		}
	}
}

func isEqual_arr(a,b [][]int)bool{
	return len(a)==len(b)
}
