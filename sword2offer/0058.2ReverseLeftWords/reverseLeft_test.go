package offer58_2

import (
	"fmt"
	"testing"
)

type question58_2 struct{
	input string
	k int
	output string
}
func Test_ReverseWord(t *testing.T){
	qs:=[]question58_2{
		{
			input: "abcdefg",
			k: 2,
			output: "cdefgab",
		},
		{
			input: "lrloseumgh",
			k: 6,
			output: "umghlrlose",
		},
	}
	for _,ex := range qs{
		k,par,as := ex.k,ex.input,ex.output
		res := reverseLeftWords(par,k)
		if res!=as{
			t.Errorf("error [input]:%v  [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}