package lt171

import (
	"fmt"
	"testing"
)

type question171 struct{
	input string
	output int
}

func Test_Pro171(t *testing.T){
	qs := []question171{
		{
			input: "A",
			output: 1,
		},
		{
			input: "AB",
			output: 28,
		},
		{
			input: "FXSHRXW",
			output: 2147483647,

		},
	}
	for _,ex := range qs{
		par,as := ex.input,ex.output
		res := titleToNumber(par)
		if res!=as{
			t.Errorf("error [input]:%v,  [out]:%v  [ans]:%v",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}