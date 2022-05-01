package lt689

import (
	"fmt"
	"testing"
)

func TestProb689(t *testing.T){
	nums := []int{1,2,1,2,6,7,5,1}
	k := 2
	res := maxSumOfThreeSubarrays(nums,k)
	fmt.Println(res)
}