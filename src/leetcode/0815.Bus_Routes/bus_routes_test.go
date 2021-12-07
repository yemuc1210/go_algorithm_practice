package leetcode

import (
	"fmt"
	"testing"
)

type question815 struct{
	par815
	ans815
}
type par815 struct{
	routes [][]int
	source int
	target int
}
type ans815 struct{
	steps int
}

func Test_BusRoutes(t *testing.T){
	qs := []question815{

		{
			par815{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6},
			ans815{2},
		},
	}
	fmt.Println("____________测试815 公交车路径（困难题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：112 ms, 在所有 Go 提交中击败了53.66%的用户
	内存消耗：22.3 MB, 在所有 Go 提交中击败了7.32%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs {
		par,as := example.par815,example.ans815

		res := numBusesToDestination(par.routes,par.source,par.target)
		if res!=as.steps{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.steps)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}