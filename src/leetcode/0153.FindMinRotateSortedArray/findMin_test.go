package lt153

import (
	"fmt"
	"testing"
)

func Test_Problem153(t *testing.T){
	input := []int{11,13,15,17}
	ans := 11
	res := findMin(input)
	if ans != res{
		t.Errorf("error res:=%v \n",res)
	}else{
		fmt.Println("pass")
	}
}