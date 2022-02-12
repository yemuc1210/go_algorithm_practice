package lt1984

import (
	"fmt"
	"testing"
)

func TestProb1984(t *testing.T) {
	nums := []int{9,4,1,7}
	res := minimumDifference(nums,2)
	ans := 2
	if res!=ans {
		t.Errorf("error ans=%v res=%v\n",ans,res)
	}else{
		fmt.Println("pass")
	}
}