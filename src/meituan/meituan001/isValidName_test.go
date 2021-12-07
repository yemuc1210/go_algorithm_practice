package meituan001

import (
	"fmt"
	"testing"
)


type question1 struct{
	par1
	ans1
}
type par1 struct{
	T int
	names []string
}
type ans1 struct{
	res []string
}

func Test_001(t *testing.T){
	qs := []question1{
		{
			par1{5,[]string{"0ooook","Hhhh666","ABCD","Meituan","6666"}},
			ans1{[]string{"Wrong","Accept","Wrong","Wrong","Wrong"}},
		},
	}

	for _,ex := range qs {
		par,as := ex.par1,ex.ans1
		res := isValidName(par.T,par.names)
		if !isEqual_arr(as.res,res){
			t.Errorf("error, input:%v   output:%v,    ans:%v\n",par,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}

func isEqual_arr(a,b []string)bool{
	if len(a) != len(b){
		return  false
	}

	for i:=0;i<len(a);i++{
		fmt.Printf("a[i]:%v,  b[i]:%v     a[i]==b[i]:%v\n",a[i],b[i],a[i]==b[i])
		if a[i]!=b[i]{
			return false
		}
	}
	return true
}