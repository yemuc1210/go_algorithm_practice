package offerS69

import (
	"fmt"
	"testing"
)

type question69 struct{
	arr []int
	ans int
}

func TestProb69(t *testing.T){
	qs := []question69{
		{
			arr: []int{0,1,0},
			ans: 1,
		},
		{
			arr: []int{0,10,5,2},
			ans: 1,
		},
		{
			arr: []int{1,3,5,4,2},
			ans: 2,
		},
	}

	for _,ex := range qs{
		par,as := ex.arr,ex.ans
		res := peakIndexInMountainArray(par)
		if res!=as{
			t.Errorf("error res=%d  ans=%d\n",res,as)
		}else{
			fmt.Println("pass")
		}
	}
}