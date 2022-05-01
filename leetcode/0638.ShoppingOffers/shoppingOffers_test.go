package lt638

import (
	"fmt"
	"testing"
)
type question638 struct{
	price,needs []int
	special [][]int
	ans int
}

func TestProblem638(t *testing.T){
	qs := []question638{
		{
			price: []int{2,5},
			special: [][]int{
				{3,0,5},{1,2,10},
			},
			needs: []int{3,2},
			ans: 14,
		},
		{
			price: []int{2,3,4},
			special: [][]int{
				{1,1,0,4},{2,2,1,9},
			},
			needs: []int{1,2,1},
			ans: 11,
		},
	}

	for _,ex := range qs {
		price,special,needs,ans := ex.price,ex.special,ex.needs,ex.ans

		res := shoppingOffers(price,special,needs)
		
		if res!=ans{
			t.Errorf("error")
		}else{
			fmt.Println("pass")
		}
	}

}