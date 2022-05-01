package lt49

import (
	"fmt"
	"testing"
)

type question49 struct{
	par49
	ans49
}
type par49 struct{
	strs []string
}
type ans49 struct{
	ans [][]string
}

func Test_Problem49(t *testing.T){
	qs := []question49{
		{
			par49{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans49{[][]string{
				{"ate","eat","tea"},{"nat","tan"},{"bat"},
			}},
		},
		
	}
	for _,ex := range qs{
		// par,_ := ex.para46,ex.ans46
		// res:=permute(par.s)
		par,_ := ex.par49,ex.ans49
		res := groupAnagrams(par.strs)
		fmt.Println(res)
	}
}