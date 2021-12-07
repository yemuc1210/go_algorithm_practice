package offer15

import (
	"fmt"
	"testing"
)
type question14 struct{
	par14
	ans14
}
type par14 struct{
	n int
}
type ans14 struct{
	out int
}
func Test_Prob14(t *testing.T){
	qs:=[]question14{
		{
			par14{2},ans14{1},
		},
		{
			par14{10},ans14{36},
		},
	}

	for _,ex :=range qs{
		par,as:=ex.par14,ex.ans14
		res:=cuttingRope(par.n)
		res1:=cuttingRope_ByMath(par.n)  //0  1.9   100% 85.87%
		res2:=cuttingRope_greedy(par.n)
		if res!=as.out || res1!=as.out || res2!=as.out{
			t.Errorf("error res=%v,ans=%v\n",res,as.out)
		}else{
			fmt.Printf("pass\n")
		}
	}
}