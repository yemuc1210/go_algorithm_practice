package leetcode

import (
	"fmt"
	"testing"
)



type question_daily1 struct{
	input uint32
	output int
}
var qs = []question_daily1{
	{
		input: uint32(00000000000000000000000000001011),
		output: 3,
	},
	{
		input: 00000000000000000000000010000000,
		output: 1,
	},
	// {
	// 	input: uint32(11111111111111111111111111111110),
	// 	output: 31,
	// },
}


func Test_HamingWeight(t *testing.T){
	for _,ques :=  range qs {
		in,out := ques.input,ques.output

		res := hammingWeight(in)

		if res != out {
			t.Errorf("error\n")
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",in,out)
		}
	}
}