package leetcode

import (
	"fmt"
	"testing"
)


type question_daily2 struct{
	point [][]int
	count int
}

func Test_Max_Points(t *testing.T){
	qs := []question_daily2{
		{
			point: [][]int{
				{1,1},{3,2},{5,3},{4,1},{2,3},{1,4},
			},
			count: 4,
		},
		{
			point: [][]int{
				{1,1},{2,2},{3,3},
			},
			count: 3,
		},{
			point: [][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}},
			count: 4,
		},
		{
			point: [][]int{
				{0,0},{1,1},
			},
			count: 2,
		},
		{
			point:[][]int{
				{0,0},{1,1},{2,2},
			},
			count: 3,
		},
		{
			point: [][]int{{0,0},{2,2},{-1,-1}},
			count: 3,
		},
	}
	fmt.Println("____________daily2 max point in line____________")
	for _,example := range qs {
		par,as := example.point,example.count

		res := maxPoints(par)
		if res != as {
			t.Errorf("error [input]:%v     [output]:%v    [ans]:%v\n",par,res,as)
		}else{
			fmt.Printf("pass [input]:%v     [output]:%v\n",par,res)
		}
	}
}