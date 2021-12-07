package leetcode

import (
	"fmt"
	"testing"
)

type question67 struct{
	par67
	ans67
}

type par67 struct{
	l1 string
	l2 string
}
type ans67 struct{
	output string
}

func Test_Add_Binary(t *testing.T){
	qs:=[]question67{
		{
			par67{l1:"1010",l2:"1011"},
			ans67{"10101"},
		},
		{
			par67{l1:"11",l2:"1"},
			ans67{"100"},
		},

	}

	fmt.Println("____________测试67 二进制加法（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了27.38%的用户
	内存消耗：2.7 MB, 在所有 Go 提交中击败了11.77%的用户`)
	for _,example := range qs {
		par,as := example.par67,example.ans67

		res := addBinary(par.l1,par.l2)
		if res != as.output{
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}