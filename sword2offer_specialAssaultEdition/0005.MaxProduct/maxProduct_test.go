package offerS5

import (
	"fmt"
	"testing"
)

type questionOfferS5 struct{
	words []string
	ans int
}

func TestProblem5(t *testing.T){
	qs := []questionOfferS5{
		{
			words: []string{"abcw","baz","foo","bar","fxyz","abcdef"},
			ans: 16,
		},
		{
			words: []string{"a","ab","abc","d","cd","bcd","abcd"},
			ans: 4,
		},
		{
			words: []string{"a","aa","aaa","aaaa"},
			ans: 0,
		},
	}

	for _,ex := range qs{
		par,ans := ex.words,ex.ans
		res := maxProduct(par)

		if res!=ans{
			t.Errorf("error ans=%v   res=%v\n",ans,res)
		}else{
			fmt.Println("pass")
		}
	}
}