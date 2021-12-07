package offer07

import (
	"fmt"
	"testing"
)

type question07 struct{
	par7
	ans7
}
type par7 struct{
	pre,in []int
}
type ans7 struct{
	out []int
}
func Test_P7(t *testing.T){
	fmt.Println("pass 4ms,4.2MB   94.85%   64.10% ")
}