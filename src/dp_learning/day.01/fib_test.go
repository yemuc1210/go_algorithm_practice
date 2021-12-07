package dpDay01

import (
	"fmt"
	"testing"
)
type question1 struct{
	input int
	output int
}

func TestFib(t *testing.T){
	qs:=[]question1{
		{
			input: 2,
			output: 1,
		},
		{
			input: 3,
			output: 2,
		},
		{
			input: 4,
			output: 3,
		},
	}
	for _,ex:=range qs{
		par,as := ex.input,ex.output
		res := fib(par)
		res1 := fib1(par)
		if res!= as && res1!=as{
			t.Errorf("error [input]:%v,[output]:%v,ans=%v\n",par,res,as)
		}
		fmt.Println("pass")
	}
}