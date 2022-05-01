package NC32

import (
	"fmt"
	"testing"
)

func TestNC32(t *testing.T){
	res := sqrt(2)
	ans := 1 
	if res!=ans{
		t.Errorf("error ans=%v  res=%v\n",ans,res)
	}else{
		fmt.Println("pass")
	}
}