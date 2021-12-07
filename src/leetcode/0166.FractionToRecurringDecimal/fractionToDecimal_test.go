package lt166

import (
	"fmt"
	"testing"
)

func TestProb166(t *testing.T){
	a,b,ans := 1,2,"0.5"

	res := fractionToDecimal(a,b)
	if res != ans{
		t.Errorf("error res=%v  ans=%v\n",res,ans)
	}else{
		fmt.Println("pass")
	}
}