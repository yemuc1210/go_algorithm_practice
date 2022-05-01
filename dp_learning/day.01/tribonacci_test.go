package dpDay01

import (
	"fmt"
	"testing"
)
type question2 struct{
	input int
	output int
}

func TestTribonacci(t *testing.T){
	qs:=[]question1{
		{
			input: 4,
			output: 4,
		},
		{
			input: 2,
			output: 1,
		},
		{
			input: 25,
			output: 1389537,
		},
	}
	for _,ex:=range qs{
		par,as := ex.input,ex.output
		res := tribonacci(par)

		if res!= as {
			t.Errorf("error [input]:%v,[output]:%v,ans=%v\n",par,res,as)
		}
		fmt.Println("pass")
	}
}