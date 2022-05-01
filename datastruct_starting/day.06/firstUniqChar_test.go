package day06

import (
	"fmt"
	"testing"
)

type question struct{
	s string
	ans int
}
func Test_FirstUniqChar(t *testing.T){
	qs:=[]question{
		{
			s: "leetcode",
			ans: 0,
		},
		{
			s: "loveleetcode",
			ans: 2,
		},
		{
			s: "aabbccddee",
			ans: -1,
		},
	}
	for _,ex := range qs{
		par,as := ex.s,ex.ans
		res:=firstUniqChar(par)
		if res!=as{
			t.Errorf("error [input]:%v    [output]:%v   [ans]:%v\n",par,res,as)
		}else{
			fmt.Printf("pass")
		}
	}
}