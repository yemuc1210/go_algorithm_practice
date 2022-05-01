package offer58

import (
	"fmt"
	"testing"
)

type question58 struct{
	input string
	output string
}
func Test_ReverseWord(t *testing.T){
	qs:=[]question58{
		{
			input: "the sky is blue",
			output: "blue is sky the",
		},
		{
			input: "    hello world!",
			output: "world! hello",
		},
		{
			input: "a good   example",
			output: "example good a",
		},
		{
			input: "",
			output: "",
		},
		{
			input: " ",
			output: "",
		},
	}
	for _,ex := range qs{
		par,as := ex.input,ex.output
		res := reverseWords(par)
		if res!=as{
			t.Errorf("error [input]:%v  [out]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}