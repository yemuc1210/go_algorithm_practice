package offer51

import (
	"fmt"
	"testing"
)

type question51 struct{
	nums []int
	res int
}

func Test_ReversePairs(t *testing.T){
	qs:=[]question51{
		{
			nums: []int{7,5,6,4},
			res: 5,
		},
		{
			nums: []int{50, 10, 20, 30, 70, 40, 80, 60},
			res: 7,
		},
	}
	for _,ex := range qs{
		par,as := ex.nums,ex.res
		res := reversePairs(par)
		if res!= as{
			t.Errorf("error [input]:%v [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}