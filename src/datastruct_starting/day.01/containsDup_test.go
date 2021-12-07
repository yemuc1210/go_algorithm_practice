package dsday01

import (
	"fmt"
	"testing"
)

type question1 struct{
	input []int
	output bool
}
func TestContainDup(t *testing.T){
	qs:=[]question1{
		{input: []int{1,2,3,1},output: true},
		{input: []int{1,2,3,4},output: false},
		{input: []int{1,1,1,3,3,4,3,2,4,2},output: true},
	}
	for _,ex:=range qs{
		par,as:=ex.input,ex.output
		res:=containsDuplicate(par)
		res1:=containsDuplicate1(par)
		if res!=as || res1!=as{
			t.Errorf("error\n")
		}else{
			fmt.Println("pass")
		}
	}
}