package lt260

import (
	"fmt"
	"sort"
	"testing"
)

type question260 struct{
	nums []int
	ans []int
}

func TestProb260(t *testing.T){
	qs := []question260{
		{
			nums: []int{1,2,1,3,2,5},
			ans: []int{3,5},
		},
	}

	for _,ex := range qs{
		nums,ans :=ex.nums,ex.ans
		res := singleNumber(nums)
		if !isEqual(res,ans) {
			t.Errorf("error ans=%v   res=%v\n",ans,res)
		}else{
			fmt.Println("pass")
		}

	}
}

func isEqual(a,b []int)bool{
	if len(a) != len(b){
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i:=0;i<len(a);i++{
		if a[i] !=b[i]{
			return false
		}
	}
	return true
}