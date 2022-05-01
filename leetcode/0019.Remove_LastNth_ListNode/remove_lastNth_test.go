package leetcode

import (
	"fmt"
	"go_practice/structs"
	"testing"
)

type question19 struct {
	par19
	ans19
}

type par19 struct {
	input *ListNode
	n int
}
type ans19 struct {
	output *ListNode
}

func Test_Q83(t *testing.T){
	qs := []question19{
		{
			par19{structs.Ints2List([]int{1,1,2}),2},
			ans19{structs.Ints2List([]int{1,2})},
		},
		{
			par19{structs.Ints2List([]int{1,1,2,3,3}),2},
			ans19{structs.Ints2List([]int{1,1,2,3})},
		},
	}

	fmt.Println("____________测试19 删除链表倒数第N个元素（中等题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了10%的用户
	内存消耗：2.2 MB, 在所有 Go 提交中击败了100%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs {
		par,as := example.par19,example.ans19

		res := removeNthFromEnd(par.input,par.n)
		if isEqual_list(res,as.output){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}

//判断是否相同，根据数字比较
func isEqual_list(l1 *ListNode,l2 *ListNode)bool{
	if l1==nil && l2==nil{    //两个链表等长，且都走到了最后
		return true  
	}
	if l1.Val == l2.Val{
		isEqual_list(l1.Next,l2.Next)
	}
	return false
}