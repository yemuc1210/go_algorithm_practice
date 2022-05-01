package leetcode

import (
	"fmt"
	// "strconv"
	"go_practice/structs"
	"testing"
)

type questionOffer37 struct{
	paro37
	anso37
}
type paro37 struct{
	input []int
}
type anso37 struct{
	output []int
}

func Test_Serialize(t *testing.T){
	qs := []questionOffer37{
		{
			paro37{[]int{
				1,2,3,structs.NULL,structs.NULL,4,5,
			}},
			anso37{[]int{
				1,2,3,structs.NULL,structs.NULL,4,5,
			}},
		},
	}
	for _,example := range qs{
		par,as:=example.paro37,example.anso37

		root := structs.Ints2TreeNode(par.input)
		res:=serialize(root)
		fmt.Printf("pass,%s,%v",res,as.output)

	}
}

