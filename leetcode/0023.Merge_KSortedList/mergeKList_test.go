package leetcode

import (
	"fmt"
	"go_practice/structs"
	"testing"
)


type question23 struct {
	para23
	ans23
}

// para 是参数
// one 代表第一个参数
type para23 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans23 struct {
	one []int
}

func Test_Problem23(t *testing.T) {

	qs := []question23{

		// {
		// 	para23{[][]int{}},
		// 	ans23{[]int{}},
		// },

		{
			para23{[][]int{
				{1},
				{1},
			}},
			ans23{[]int{1, 1}},
		},

		{
			para23{[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			para23{[][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			para23{[][]int{
				{1},
				{9, 9, 9, 9, 9},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para23{[][]int{
				{9, 9, 9, 9, 9},
				{1},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para23{[][]int{
				{2, 3, 4},
				{4, 5, 6},
			}},
			ans23{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			para23{[][]int{
				{1, 3, 8},
				{1, 7},
			}},
			ans23{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Println("____________测试23 合并k个有序链表（困难题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：8 ms, 在所有 Go 提交中击败了94.56%的用户
	内存消耗：6 MB, 在所有 Go 提交中击败了32.94%的用户
	递归改成迭代，效率会更高
	`)
	for _,example := range qs {
		par,as := example.para23,example.ans23
		var ls []*ListNode
		for _, qq := range par.one {
			ls = append(ls, structs.Ints2List(qq))
		}
		res := mergeKLists(ls)
		// res := mergeTwoLists(par.l1,par.l2)
		ansList := structs.Ints2List(as.one)
		if isEqual(res,ansList){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as.one)
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