package leetcode

import (
	"fmt"
	"testing"
)

type question17_10 struct{
	input []int
	output int
}
func Test_Majority(t *testing.T){
	qs:=[]question17_10{
		{
			input: []int{1,2,5,9,5,9,5,5,5},
			output: 5,
		},
		{
			input: []int{3,2},
			output: -1,
		},
		{
			input: []int{2,2,1,1,1,2,2},
			output: 2,
		},
	}
	for _,ex:=range qs{
		par,as := ex.input,ex.output
		res:=majorityElement(par)
		if res!=as{
			t.Errorf("error")
		}else{
			fmt.Println("pass")
		}
	}
}