package interview01

import (
	"fmt"
	"testing"
)

type question struct{
	input string
	output string
}

func Test_CompressString(t *testing.T){
	qs := []question{
		{
			input: "aabcccccaaa",
			output: "a2b1c5a3",
		},
		{
			input: "abbccd",
			output: "abbccd",
		},
	}

	for _,ex := range qs{
		par,as := ex.input,ex.output

		res := compressString(par)
		if res!=as{
			t.Errorf("input:%v  ans:%v   output:%v\n",par,as,res)
		}else{
			fmt.Println("pass")
		}
	}
}