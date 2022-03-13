package greedy
import (
	"testing"
	"fmt"
)
func TestProb455(t *testing.T) {
	g := []int{1,2}
	s := []int{1,2,3}
	res := findContentChildren(g,s)
	fmt.Println(res)
}