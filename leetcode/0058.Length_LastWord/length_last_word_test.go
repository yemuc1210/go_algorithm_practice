package leetcode

import (
	"fmt"
	"testing"
)

type question58 struct{
	par58
	ans58
}

type par58 struct{
	input string
}
type ans58 struct{
	output int
}

func Test_LenOfLastWord(t *testing.T){
	qs := []question58{
		{
			par58{"Hello world"},
			ans58{5},
		},
		{
			par58{" "},
			ans58{0},
		},
	}
	fmt.Println("____________测试58 最后一个单词长度（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了74.11%的用户`)
	for _,example := range qs {
		par,as := example.par58,example.ans58

		res := lengthOfLastWord(par.input)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par.input,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par.input,res)
		}
	}
}