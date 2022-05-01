package leetcode

import (
	"fmt"
	"testing"
)




type question14 struct{
	par14
	ans14
}

type par14 struct{
	input []string
}

type ans14 struct{
	output string
}
func Test_LongestCommonPrefix(t *testing.T){
	qs := []question14{
		{
			par14{[]string{"flower","flow","flight"}},
			ans14{"fl"},
		},
		{
			par14{[]string{"dog","racecar","car"}},
			ans14{""},
		},
	}

	fn0(qs, t)


}

func fn0(qs []question14, t *testing.T) {
	fmt.Println("____________测试14 最长公共前缀（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2.4 MB, 在所有 Go 提交中击败了11.77%的用户`)
	for _, example := range qs {
		par, as := example.par14, example.ans14

		res := longestCommonPrefix(par.input)
		if res != as.output {
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n", par.input, res, as.output)
		} else {
			fmt.Printf("pass [input]:%v     [output]:%v\n", par.input, res)
		}
	}
}