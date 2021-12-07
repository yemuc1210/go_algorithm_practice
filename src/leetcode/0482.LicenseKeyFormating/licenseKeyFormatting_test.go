package lt482

import (
	"fmt"
	"testing"
)

func TestProb482(t *testing.T) {
	s, k := "2-5g-3-J", 2
	ans := "2-5G-3J"
	res := licenseKeyFormatting(s, k)
	if ans != res {
		t.Errorf("error ans=%v  res=%v\n", ans, res)
	} else {
		fmt.Println("pass")
	}
}
