package offer57

import (
	"fmt"
	"testing"
)

type question57 struct{
	nums []int
	target int
	res []int
}

func Test_TwoSum(t *testing.T){
	qs:=[]question57{
		{
			nums: []int{2,7,11,15},
			target: 9,
			res: []int{7,2},
		},
		{
			nums: []int{10,26,30,31,47,69},
			target: 40,
			res: []int{30,10},
		},
	}
	for _,ex := range qs{
		nums,target,as := ex.nums,ex.target,ex.res
		res := twoSum(nums,target)
		if !isEqual_arr(res,as){
			t.Errorf("error [input]:%v %v [out]:%v   [ans]:%v\n",nums,target,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}
func isEqual_arr(a,b []int)bool{
	if len(a)!=len(b){
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			return false
		}
	}
	return true
}