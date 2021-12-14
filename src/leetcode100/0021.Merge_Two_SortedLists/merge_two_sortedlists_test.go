package leetcode

import (
	"fmt"
	"structs"
	"testing"
)

// type ListNode = structs.ListNode
type question21 struct{
	par21
	ans21
}

type par21 struct{
	l1 *ListNode
	l2 *ListNode
}
type ans21 struct{
	output *ListNode
}

func Test_Merge_Two_SortedLists(t *testing.T){
	qs := []question21{
		{
			par21{l1: structs.Ints2List([]int{1,2,4}),l2:structs.Ints2List([]int{1,3,4})},
			ans21{structs.Ints2List([]int{1,1,2,3,4,4})},
		},
		{
			par21{l1: structs.Ints2List([]int{}),l2:structs.Ints2List([]int{0})},
			ans21{structs.Ints2List([]int{0})},
		},
	}
	
	fmt.Println("____________测试21 合并两个有序链表（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了66.76%的用户
	内存消耗：2.5 MB, 在所有 Go 提交中击败了27.22%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs {
		par,as := example.par21,example.ans21

		res := mergeTwoLists(par.l1,par.l2)
		if isEqual(res,as.output){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.output)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}

//判断是否相同，根据数字比较
func isEqual(l1 *ListNode,l2 *ListNode)bool{
	if l1==nil && l2==nil{    //两个链表等长，且都走到了最后
		return true  
	}
	if l1.Val == l2.Val{
		isEqual(l1.Next,l2.Next)
	}
	return false
}