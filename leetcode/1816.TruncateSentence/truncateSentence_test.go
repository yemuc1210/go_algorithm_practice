package lt1816

import (
	"testing"
	"fmt"
)

func TestPron1816(t *testing.T){
	s := "hello how are you Contestant"
	k := 4
	ans := "hello how are you"
	res := truncateSentence(s,k)
	if ans != res {
		fmt.Println("error")
	}else{
		fmt.Println("pass")
	}
}