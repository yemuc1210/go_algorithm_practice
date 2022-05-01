package dsday01

import (
	"fmt"
	"testing"
)

type question53 struct{
	par53
	ans53
}

type par53 struct{
	input []int
}
type ans53 struct{
	output int
}

func Test_MaxSubArr(t *testing.T){
	qs := []question53{
		{
			par53{[]int{-2,1,-3,4,-1,2,1,-5,4}},
			ans53{6},
		},
		{
			par53{[]int{1}},
			ans53{1},
		},
		{
			par53{[]int{-10000}},
			ans53{-10000},
		},
	}
	fmt.Println("____________测试53 删除元素（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：8 ms, 在所有 Go 提交中击败了57.27%的用户
	内存消耗：3.4 MB, 在所有 Go 提交中击败了17.49%的用户`)
	for _,example := range qs {
		par,as := example.par53,example.ans53

		res := maxSubArray(par.input)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par.input,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par.input,res)
		}
	}
}