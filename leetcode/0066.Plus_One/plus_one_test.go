package leetcode

import (
	"fmt"
	"testing"
)

type question66 struct{
	par66
	ans66
}

type par66 struct{
	input []int
}
type ans66 struct{
	output []int
}

func Test_PlusOne(t *testing.T){

	qs:=[]question66{
		{
			par66{[]int{1,2,3}},
			ans66{[]int{1,2,4}},
		},
		{
			par66{[]int{4,3,2,1}},
			ans66{[]int{4,3,2,2}},
		},
		{
			par66{[]int{0}},
			ans66{[]int{1}},
		},
		{
			par66{[]int{9}},
			ans66{[]int{1,0}},
		},
	}

	fmt.Println("____________测试66 最后一个单词长度（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2 MB, 在所有 Go 提交中击败了95.89%的用户`)
	for _,example := range qs {
		par,as := example.par66,example.ans66

		res := plusOne(par.input)
		if !isEqual(res,as.output){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par.input,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par.input,res)
		}
	}
}

func isEqual(a []int,b []int)bool{
	if len(a) != len(b){
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i] != b[i]{
			return false
		}
	}
	return true
	//或者转化成字符串，然后在变成数字,应该strConv.atoi
}