package lt1743

import (
	"fmt"
	"testing"
)

/*
输入：adjacentPairs = [[2,1],[3,4],[3,2]]
输出：[1,2,3,4]
*/
type question1743 struct{
	adjacentPairs [][]int
	ans []int
}

func Test_Problem1743(t *testing.T){
	qs :=[]question1743{
		{
			adjacentPairs: [][]int{{2,1},{3,4},{3,2}},
			ans: []int{1,2,3,4},
		},
	}
	for _,ex := range qs{
		par,as := ex.adjacentPairs,ex.ans
		res := restoreArray(par)
		if isSameArr(res,as){
			fmt.Println("pass")

		}else{
			t.Errorf("error")
		}
	}
}
func isSameArr(a,b []int)bool{
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