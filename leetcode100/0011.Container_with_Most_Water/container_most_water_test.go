package leetcode

import (
	"fmt"
	"testing"
)

type question11 struct{
	para11
	ans11
}

type para11 struct{
	height []int
}

type ans11 struct{
	res int
}


func Test_Find_Most_Water(t *testing.T){
	
qs:= []question11{
	{
		para11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		ans11{49},
	},

	{
		para11{[]int{1, 1}},
		ans11{1},
	},
}
	fmt.Println("_________测试11 最大容器___________")
	for _,qus := range qs {
		par,ans := qus.para11,qus.ans11

		res := find_most_container(par.height)

		if res != ans.res {
			t.Errorf("error [input]:%v       [output]:%v\n",par.height,res)
		}
		fmt.Printf("[input]:%v    [output]:%v\n",par.height,res)
	}
}