package lt434

import (
	"fmt"
	"testing"
)

func TestPro434(t *testing.T){
	par,ans := "Of all the gin joints in all the towns in all the world,   ",13
	res := countSegments(par)

	if res!= ans{
		t.Errorf("error\n")
	}else{
		fmt.Println("pass")
	}
}