package leetcode

import (
	"fmt"
	"testing"
)



type question752 struct{
	par752
	ans752
}

type par752 struct{
	deadends []string
	target string
}

type ans752 struct{
	output int
}

func Test_Open_Lock(t *testing.T){
	qs := []question752{
		{
			par752{deadends:[]string{"0201","0101","0102","1212","2002"},target: "0202"},
			ans752{6},
		},
		{
			par752{deadends :[]string{"8888"}, target :"0009"},
			ans752{1},
		},
		{
			par752{deadends:[]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, target:"8888"},
			ans752{-1},
		},
		{
			par752{deadends:[]string{"0000"}, target :"8888"},
			ans752{-1},
		},

	}

	fmt.Println("____________daily3 open the lock____________")

	for _,example := range qs {
		par,as := example.par752,example.ans752

		res :=Sol_1_2(par.deadends,par.target)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}