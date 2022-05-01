package offer20

import (
	"fmt"
	"testing"
)

func Test_IsInteger(t *testing.T){
	qs:=[]string{"12e","1a3","12e123   ","11e-13","123"}
	for _,ex:=range qs{
		fmt.Printf("ex=%v,res=%v\n",ex,isInteger(ex))
	}
}
func Test_IsDeli(t *testing.T){
	qs:=[]string{".1","1.21e","1.2    ","1.2.23",".123","1."}
	for _,ex:=range qs{
		fmt.Printf("ex=%v,res=%v\n",ex,isDeli(ex))
	}
}
type question20 struct{
	par20
	ans20
}
type par20 struct{
	s string
}
type ans20 struct{
	out bool
}
func Test_IsNum(t *testing.T){
	qs:=[]question20{
		{
			par20{"+100"},ans20{true},
		},
		{par20{"5e2"},ans20{true}},
		{par20{"-123"},ans20{true}},
		{par20{"3.1415"},ans20{true}},
		{par20{"-1E-14  "},ans20{true}},
		{par20{"0123"},ans20{true}},
		{par20{"12e"},ans20{false}},
		{par20{"1s3.12"},ans20{false}},
		{par20{"1.2.3"},ans20{false}},
		{par20{"+-5"},ans20{false}},
		{par20{"12e+5.4"},ans20{false}},
		{par20{" "},ans20{false}},
		{par20{"e9"},ans20{false}},
		{par20{".2"},ans20{true}},
		{par20{"3. "},ans20{true}},
		{par20{"959440.94f"},ans20{false}},
		{par20{".."},ans20{false}},
		{par20{". "},ans20{false}},
	}
	for _,ex:=range qs{
		par,as:=ex.par20,ex.ans20
		res:=isNumber(par.s)
		if res!=as.out{
			t.Errorf("error par=%v, res=%v,ans=%v\n",par.s,res,as.out)
		}else{
			fmt.Printf("pass par=%v, res=%v,ans=%v\n",par.s,res,as.out)
		}
	}
}