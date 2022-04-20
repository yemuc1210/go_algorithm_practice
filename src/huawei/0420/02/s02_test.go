package s02

import (
	"fmt"
	"testing"
)

func TestGetInput(t *testing.T) {
	root,path,sub := getInput()
	fmt.Println(root)
	fmt.Println(path)
	fmt.Println(sub)
}

func TestSolve(t *testing.T) {
	root :=[]int{1,1,2,0,0,4,5}
	path:="/1/1"
	sub:=[]int{5,3,0}
	solve(root,path,sub)
}