package leetcode

import (
	"fmt"
	"testing"
)

type question69 struct{
	par69
	ans69
}

type par69 struct{
	input int
}
type ans69 struct{
	output int
}

func Test_My_Sqrt(t *testing.T){
	qs := []question69{
		{
			par69{4},
			ans69{2},
		},
		{
			par69{8},
			ans69{2},
		},
	}

	fmt.Println("____________测试69 求平方根（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了10.00%的用户
	内存消耗：2.1 MB, 在所有 Go 提交中击败了99.94%的用户`)
	for _,example := range qs {
		par,as := example.par69,example.ans69

		res := mySqrt(par.input)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}