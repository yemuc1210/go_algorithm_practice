package leetcode

import (
	"fmt"
	"structs"
	"testing"
)

type question83 struct {
	par83
	ans83
}

type par83 struct {
	input *ListNode
}
type ans83 struct {
	output *ListNode
}

func Test_Q83(t *testing.T){
	qs := []question83{
		{
			par83{structs.Ints2List([]int{1,1,2})},
			ans83{structs.Ints2List([]int{1,2})},
		},
		{
			par83{structs.Ints2List([]int{1,1,2,3,3})},
			ans83{structs.Ints2List([]int{1,2,3})},
		},
	}

	fmt.Println("____________测试83 去除有序链表重复元素（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：4 ms, 在所有 Go 提交中击败了85.74%的用户
	内存消耗：3.1 MB, 在所有 Go 提交中击败了59.30%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs {
		par,as := example.par83,example.ans83

		res := deleteDuplicates(par.input)
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