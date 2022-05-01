package leetcode

import (
	"fmt"
	"testing"
)

type question70 struct{
	par70
	ans70
}
type par70 struct{
	input int
}
type ans70 struct{
	output int
}

func Test_Climb_Stairs(t *testing.T){
	qs := []question70{
		{
			par70{2},ans70{2},
		},
		{
			par70{3},ans70{3},
		},
	}

	
	fmt.Println("____________测试70 爬台阶（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了10.00%的用户
	内存消耗：1.9 MB, 在所有 Go 提交中击败了78.76%的用户`)
	for _,example := range qs {
		par,as := example.par70,example.ans70

		res := climbStairs(par.input)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}