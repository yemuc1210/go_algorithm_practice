package lt322

import (
	"fmt"
	"testing"
)

type question322 struct{
	coins []int
	amount int
	ans int
}
func TestCoinChange(t *testing.T) {
	qs := []question322{
		{
			coins: []int{1,2,5},
			amount: 11,
			ans: 3,
		},
		{
			coins: []int{2},
			amount: 3,
			ans: -1,
		},
		{
			coins: []int{1},
			amount: 0,
			ans: 0,
		},
		{
			coins: []int{1},
			amount: 1,
			ans: 1,
		},
	}

	for _,ex := range qs {
		coins,amount,ans := ex.coins,ex.amount,ex.ans

		res := coinChange(coins,amount)

		if res!= ans {
			t.Errorf("error ans=%d  res=%d\n",ans,res)
		}else{
			fmt.Println("pass")
		}
	}
}