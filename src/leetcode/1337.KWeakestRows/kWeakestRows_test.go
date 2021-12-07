package lt1337

import (
	"fmt"
	"testing"
)

type question1337 struct{
	mat [][]int
	k int
	ans []int
}

func Test_Pro1337(t *testing.T){
	qs:=[]question1337{
		{
			mat: [][]int{
				{1,1,0,0,0},
					{1,1,1,1,0},
					{1,0,0,0,0},
					{1,1,0,0,0},
					{1,1,1,1,1},
			},
			k: 3,
			ans: []int{2,0,3},
		},
	}
	for _,ex := range qs{
		mat,k,ans := ex.mat,ex.k,ex.ans

		res := kWeakestRows(mat,k)
		fmt.Println(res)
		fmt.Println(ans)
	}
}