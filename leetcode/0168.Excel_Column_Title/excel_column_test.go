package leetcode

import (
	"fmt"
	"testing"
)


type question168 struct {
	para168
	ans168
}

// para 是参数
// one 代表第一个参数
type para168 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans168 struct {
	one string
}

func Test_Problem168(t *testing.T) {

	qs := []question168{

		{
			para168{1},
			ans168{"A"},
		},

		{
			para168{28},
			ans168{"AB"},
		},

		{
			para168{701},
			ans168{"ZY"},
		},

		{
			para168{10011},
			ans168{"NUA"},
		},

		{
			para168{999},
			ans168{"ALK"},
		},

		{
			para168{681},
			ans168{"ZE"},
		},
	}
	fmt.Println("____________测试168 EXCEL表格列名称（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户`)
	for _,example := range qs {
		par,as := example.para168,example.ans168

		res := convertToTitle(par.n)
		if res != as.one{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.one)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}