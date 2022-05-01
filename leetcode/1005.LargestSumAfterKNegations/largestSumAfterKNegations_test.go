package lt1005

import (
	"fmt"
	"testing"
)

func TestLt1005(t *testing.T){
	nums := []int{2,-3,-1,5,-4}
	ans := 13
	res := largestSumAfterKNegations(nums,2)
	if res!=ans{
		t.Errorf("error ans=%v  res=%v\n",ans,res)
	}else{
		fmt.Println("pass")
	}
}