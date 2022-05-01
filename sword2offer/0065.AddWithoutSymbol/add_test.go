package offer65

import (
	"fmt"
	"testing"
)

type question65 struct{
	a,b int
	ans int
}

func Test_Add(t *testing.T){
	qs :=[]question65{
		{
			a: 1,
			b: 1,
			ans: 2,
		},
		{
			a: 2,
			b: 9,
			ans: 11,
		},
	}
	for _,ex := range qs{
		a,b,as:=ex.a,ex.b,ex.ans
		res:=add(a,b)
		if res!=as{
			t.Errorf("error [input]:%v  %v  [out]:%v   [ans]:%v\n",a,b,res,as)
		}else{
			fmt.Println("pass")
		}
	}

}