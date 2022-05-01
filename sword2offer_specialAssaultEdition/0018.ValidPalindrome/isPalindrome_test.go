package offerS18

import (
	"fmt"
	"testing"
)

type quetin18 struct{
	s string
	ans bool
}

func TestProb18(t *testing.T){
	qs := []quetin18{
		{
			s: "A man, a plan, a canal: Panama",
			ans: true,
		},
	}
	for _,ex := range qs{
		par,as := ex.s,ex.ans
		res := isPalindrome(par)
		if as !=res{
			t.Errorf("error\n")
		}else{
			fmt.Println("pass")
		}
	}
}