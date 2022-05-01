package leetcode

import (
	"fmt"
	"testing"
)

type question28 struct{
	par28
	ans28
}
type par28 struct{
	haystack string
	needle string
}

type ans28 struct{
	output int
}

func Test_strStr(t *testing.T){
	qs := []question28{
		{
			par28{haystack:"hello",needle: "ll"},
			ans28{2},
		},
		{
			par28{haystack:"aaaaa", needle:"bba"},
			ans28{-1},
		},
		{
			par28{haystack:"",needle:""},
			ans28{0},
		},
	}

	fmt.Println("____________测试27 删除元素（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100%的用户
	内存消耗：2.2 MB, 在所有 Go 提交中击败了99.93%的用户`)
	for _,example := range qs {
		par,as := example.par28,example.ans28

		res := strStr(par.haystack,par.needle)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}