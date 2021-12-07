package offer06

import (
	"fmt"
	"structs"
	"testing"
)

type question06 struct{
	par06
	ans06
}

type par06 struct{
	l1 *ListNode
}
type ans06 struct{
	output []int
}


func Test_PrintListReverse(t *testing.T){
	qs := []question06{
		{
			par06{l1: structs.Ints2List([]int{1,3,2})},
			ans06{[]int{2,3,1}},
		},
	}
	fmt.Println("____________offer06 从尾到头打印链表（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了60.84%的用户
	内存消耗：3.5 MB, 在所有 Go 提交中击败了49.97%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs {
		par,as := example.par06,example.ans06

		res := reversePrint(par.l1)
		res1:=reversePrint1(par.l1)
		if !isEqual_arr(res,as.output) && !isEqual_arr(res1,as.output){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
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