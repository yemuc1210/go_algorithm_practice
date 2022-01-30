package NC41

import (
	"fmt"
	"testing"
)

func TestProbNC41(t *testing.T){
	arr := []int{2,3,4,5}
	ans := 4
	res := maxLength(arr)
	if res!=ans {
		t.Errorf("error  ans=%v  res=%d\n",ans,res)
	}else{
		fmt.Println("pass")
	}
}