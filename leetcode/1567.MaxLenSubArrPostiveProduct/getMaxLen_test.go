package lt1567

import (
	"fmt"
	"testing"
)

type question1567 struct{
	nums []int
	maxProduct int
}

func Test_MaxProduct(t *testing.T){
	qs:=[]question1567{
		{
			nums: []int{1,-2,-3,4},
			maxProduct: 4,
		},
	}
	for _,ex:=range qs{
		par,as := ex.nums,ex.maxProduct
		res:=getMaxLen(par)
		if res!= as{
			t.Errorf("error [input]:%v [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}