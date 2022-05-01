package leetcode

import (
	"fmt"
	"testing"
)

type question30 struct {
	para30
	ans30
}
type para30 struct {
	s     string
	words []string
}
type ans30 struct {
	output []int
}

func Test_Find_SubStrinf(t *testing.T) {
	qs := []question30{
		{
			para30{"aaaaaaaa", []string{"aa", "aa", "aa"}},
			ans30{[]int{0, 1, 2}},
		},

		{
			para30{"barfoothefoobarman", []string{"foo", "bar"}},
			ans30{[]int{0, 9}},
		},

		{
			para30{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}},
			ans30{[]int{}},
		},

		{
			para30{"goodgoodgoodgoodgood", []string{"good"}},
			ans30{[]int{0, 4, 8, 12, 16}},
		},

		{
			para30{"barofoothefoolbarman", []string{"foo", "bar"}},
			ans30{[]int{}},
		},

		// {
		// 	para30{"bbarffoothefoobarman", []string{"foo", "bar"}},
		// 	ans30{[]int{}},
		// },

		// {
		// 	para30{"ooroodoofoodtoo", []string{"foo", "doo", "roo", "tee", "oo"}},
		// 	ans30{[]int{}},
		// },

		{
			para30{"abc", []string{"a", "b", "c"}},
			ans30{[]int{0}},
		},

		{
			para30{"a", []string{"b"}},
			ans30{[]int{}},
		},

		{
			para30{"ab", []string{"ba"}},
			ans30{[]int{}},
		},

		{
			para30{"n", []string{}},
			ans30{[]int{}},
		},
	}
	fmt.Println("____________测试30 串联所有单词的子串（困难题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：908 ms, 在所有 Go 提交中击败了7.22%的用户
	内存消耗：7.5 MB, 在所有 Go 提交中击败了18.98%的用户`)
	for _, example := range qs {
		par, as := example.para30, example.ans30
		res := findSubstring(par.s, par.words)
		if !isEqual(res, as.output) {
			t.Errorf("error [input]:%v   [output]:%v\n", par, res)
		} else {
			fmt.Printf("pass [input]:%v   [output]:%v\n", par, res)
		}
	
	}
	fmt.Println("____________测试30 串联所有单词的子串 题解思路（困难题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：3.3 MB, 在所有 Go 提交中击败了82.89%的用户`)
	for _, example := range qs {
		par, as := example.para30, example.ans30
		res := findSubstring_1(par.s, par.words)
		if !isEqual(res, as.output) {
			t.Errorf("error [input]:%v   [output]:%v\n", par, res)
		} else {
			fmt.Printf("pass [input]:%v   [output]:%v\n", par, res)
		}
	}
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
