package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}

	q := []*Node{root}
	for len(q) > 0 {
		curWidth := len(q)
		curLevelNodes := []int{}
		for i := 0; i < curWidth; i++ {
			curNode := q[0]
			q := q[1:]
			curLevelNodes = append(curLevelNodes, curNode.Val)
			// 查询curNode的Children
			for i := 0; i < len(curNode.Children); i++ {
				if curNode.Children[i] != nil {
					q = append(q, curNode.Children[i])
				}
			}
		}

		res = append(res, curLevelNodes)

	}

	return res
}
func levelOrder1(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*Node{root}
	for q != nil {
		level := []int{}
		tmp := q
		q = nil
		for _, node := range tmp {
			level = append(level, node.Val)
			q = append(q, node.Children...)
		}
		ans = append(ans, level)
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
