package offer61

import (
	"fmt"
	"testing"
)
type question61 struct{
	nums []int
	ans bool
}
func Test_IsStraight(t *testing.T){
	qs:=[]question61{
		{
			nums: []int{1,2,4,3,5},
			ans: true,
		},
		{
			nums:[]int{1,2,5,0,0},
			ans: true,
		},
	}
	for _,ex := range qs{
		nums,as := ex.nums,ex.ans
		res:=isStraight(nums)
		res1 :=isStraight1(nums)
		if res!=as || res1!=as{
			t.Errorf("error [input]:%v  [out]:%v   [ans]:%v\n",nums,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}