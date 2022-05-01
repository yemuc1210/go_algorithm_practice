package offer17

import (
	"fmt"
	"testing"
)

type question7 struct{
	n int
	out []int
}
func Test_PrintNum(t *testing.T){
	qs:=[]question7{
		
		{
			n:1,
			out: []int{1,2,3,4,5,6,7,8,9},
		},
	}
	for _,ex:=range qs{
		res:=printNumbers(ex.n)
		if !isEqual_arr(res,ex.out){
			t.Errorf("error\n")
		}else{
			fmt.Println("pass")
		}
	}
}
func isEqual_arr(a,b []int)bool{
	if len(a)!=len(b){
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			return false
		}
	}
	return true
}