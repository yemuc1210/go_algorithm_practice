package leetcode

import (
	"fmt"
	"testing"
)

type question88 struct{
	par88
	ans88
}

type par88 struct{
	nums1 []int
	nums2 []int
	m int
	n int
}

type ans88 struct{
	out []int
}

func Test_Merge_SortedArr(t *testing.T){
	qs :=[]question88{
		{
			par88{nums1: []int{1,2,3,0,0,0},m:3,nums2:[]int{2,5,6},n:3},
			ans88{[]int{1,2,2,3,5,6}},
		},
	}
	fmt.Println("____________测试88 合并有序数组（简单题）____________")
	fmt.Println(`执行结果：
	通过
	执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
	内存消耗：2.3 MB, 在所有 Go 提交中击败了99.97%的用户`)
	for _,example := range qs {
		par,as := example.par88,example.ans88

		merge(par.nums1,par.m,par.nums2,par.n)

		if !isEqual_arr(par.nums1,as.out){
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par.nums1,par.nums1,as.out)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par.nums1,par.nums1)
		}
	}
}

func isEqual_arr(a []int,b []int)bool{
	if len(a) != len(b){
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i] != b[i]{
			return false
		}
	}
	return true
}