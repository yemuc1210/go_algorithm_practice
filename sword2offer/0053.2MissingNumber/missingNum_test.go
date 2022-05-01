package offer53

import (
	"fmt"
	"testing"
)

type question52 struct{
	nums []int
	ans int
}
func Test_MissingNum(t *testing.T){
	qs:=[]question52{
		{
			nums: []int{0,1,3},
			ans: 2,
		},
		{
			nums: []int{0,1,2,3,4,5,6,7,9},
			ans: 8,
		},
		{
			nums: []int{0},
			ans: 1,
		},
	}
	for _,ex := range qs{
		par,as := ex.nums,ex.ans
		res := missingNumber(par)
		if res!= as{
			t.Errorf("error [input]:%v [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}