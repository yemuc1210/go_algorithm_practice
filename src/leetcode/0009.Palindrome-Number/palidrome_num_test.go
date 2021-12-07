package leetcode

import (
	"fmt"
	"testing"
)


type question9 struct{
	para9
	ans9
}

type para9 struct{
	number int
}

type ans9 struct{
	isPalidrome bool
}



func Test_Palidrome_Number(t *testing.T){
	qs := []question9{
		{
			para9{121},
			ans9{true},
		},
		{
			para9{-121},
			ans9{false},
		},
		{
			para9{10},
			ans9{false},
		},
		{
		para9{1534236469},
		ans9{false},
		},
		{
			para9{120},
			ans9{false},
		},
		{
			para9{321},
			ans9{false},
		},
	}

	fmt.Println("__________测试9 回文整数判断______________")
	for _,example := range qs{
		ans,par := example.ans9,example.para9

		res := palidreome_number(par.number)

		if res != ans.isPalidrome {
			t.Errorf("error [input]:%v       [output]:%v\n",par.number,res)
		}

		fmt.Printf("[input]:%v       [output]:%v\n",par.number,res)
	}

}

func Test_Palidrome_Number_by_Str(t *testing.T){
	qs := []question9{
		{
			para9{121},
			ans9{true},
		},
		{
			para9{-121},
			ans9{false},
		},
		{
			para9{10},
			ans9{false},
		},
		{
		para9{1534236469},
		ans9{false},
		},
		{
			para9{120},
			ans9{false},
		},
		{
			para9{321},
			ans9{false},
		},
	}

	fmt.Println("__________测试9 回文整数判断 str______________")
	for _,example := range qs{
		ans,par := example.ans9,example.para9

		res := palidreome_number_by_string(par.number)

		if res != ans.isPalidrome {
			t.Errorf("error [input]:%v       [output]:%v\n",par.number,res)
		}

		fmt.Printf("[input]:%v       [output]:%v\n",par.number,res)
	}

}