package offerS94

import (
	"fmt"
	"testing"
)

func TestProb94(t *testing.T){
	s := "aab"
	ans :=1
	res := minCut(s)

	fmt.Println(ans,res)
}