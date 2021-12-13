package lt807

import (
	"fmt"
	"testing"
)

func TestProb807(t *testing.T){
	grid := [][]int{
		{3,0,8,4},
		{2,4,5,7},
		{9,2,6,3},
		{0,3,1,0},
	}
	res := maxIncreaseKeepingSkyline(grid)
	ans := 35
	if res!=ans {
		t.Errorf("error ans=%v  res=%v\n",ans,res)
	}else{
		fmt.Println("pass")
	}
}