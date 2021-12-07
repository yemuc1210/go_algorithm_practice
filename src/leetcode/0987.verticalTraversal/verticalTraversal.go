package lt987

import (
	"sort"
	"structs"
)
type TreeNode = structs.TreeNode
/*
垂序遍历
*/
type node struct {
	x, y, val int
}

func verticalTraversal(root *TreeNode) [][]int {
	nodes := []*node{}
	inorder(root, 0, 0, &nodes)
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].y == nodes[j].y {
			if nodes[i].x < nodes[j].x {
				return true
			} else if nodes[i].x > nodes[j].x {
				return false
			}
			return nodes[i].val < nodes[j].val
		}
		return nodes[i].y < nodes[j].y
	})
	res, currY, currColumn := [][]int{}, nodes[0].y, []int{nodes[0].val}
	for i := 1; i < len(nodes); i++ {
		if currY == nodes[i].y {
			currColumn = append(currColumn, nodes[i].val)
		} else {
			res = append(res, currColumn)
			currColumn = []int{nodes[i].val}
			currY = nodes[i].y
		}
	}
	res = append(res, currColumn)
	return res
}

func inorder(root *TreeNode, x, y int, nodes *[]*node) {
	if root != nil {
		*nodes = append(*nodes, &node{x, y, root.Val})
		inorder(root.Left, x+1, y-1, nodes)
		inorder(root.Right, x+1, y+1, nodes)
	}
}