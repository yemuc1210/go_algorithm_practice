package lt152

import (
	"fmt"
	"testing"
)

type question152 struct{
	nums []int
	maxProduct int
}

func Test_MaxProduct(t *testing.T){
	qs:=[]question152{
		{
			nums: []int{2,3,-2,4},
			maxProduct: 6,
		},
	}
	for _,ex:=range qs{
		par,as := ex.nums,ex.maxProduct
		res:=maxProduct(par)
		if res!= as{
			t.Errorf("error [input]:%v [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}