package offerS35

import (
	"fmt"
	"testing"
)

type question35 struct{
	timePoints []string
	ans int
}

func TestPro35(t *testing.T) {
	qs := []question35{
		{
			timePoints: []string{"23:59","00:00"},
			ans: 1,
		},
		{
			timePoints: []string{"00:00","23:59","00:00"},
			ans: 0,
		},
	}

	for _,ex := range qs{
		par,as := ex.timePoints,ex.ans
		res := findMinDifference(par)
		if res!=as{
			t.Errorf("error")
		}else{
			fmt.Println("pass")
		}
	}
}