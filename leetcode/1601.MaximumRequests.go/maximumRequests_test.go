package lt1601

import (
	"fmt"
	"testing"
)

func TestProb1601(t *testing.T) {
	requests := [][]int{
		{0,1},{1,0},{1,2},{2,0},{3,4},
	}
	res := maximumRequests(5,requests)
	fmt.Println(res)
}