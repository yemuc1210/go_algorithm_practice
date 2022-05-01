package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63
var m = make(map[int]int)
var sameNodes = make(map[int]map[*TreeNode]int)

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	// 如果当前值是重复出现的
	if m[ints[0]] != 0 {
		// 存入map
		sameNodes[ints[0]][root]++
	}

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		// 如果当前值是重复出现的
		if m[ints[i]] != 0 {
			// 存入map
			sameNodes[ints[i]][root]++
		}
		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func main() {
	//获得输入

	nodes := []int{1, 2, 3, 1, NULL, 2, NULL, NULL, NULL, NULL, NULL, 1, NULL, NULL, NULL}
	for _, v := range nodes {
		if v != NULL {
			m[v]++
		}
	}
	// 处理相同的节点
	//sameNodes  存了重复的值，是可疑的
	root := Ints2TreeNode(nodes)

	fmt.Println(Tree2ints(root))

	var cnt int
	// 求解
	// 思路为：比较两个子树是否想等
	// 另一个关键算法是得到相等的节点
	for _, v := range sameNodes {
		// k值出现多次  v:map[*TreeNode]int
		// 取一个节
		var nodes []*TreeNode
		for k1 := range v {
			nodes = append(nodes, k1)
		}
		oneNode := nodes[0]
		nodes = nodes[1:]
		for i := range nodes {
			if isSameTree(oneNode, nodes[i]) {
				cnt++
				// 输出

			}
		}
		if cnt != 0 {
			fmt.Println(Tree2ints(oneNode))
			cnt = 0
		}
	}
	// over
}

func isSameTree(root1, root2 *TreeNode) bool {
	return root1.Val == root2.Val && isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
}

// GetTargetNode 返回 Val = target 的 TreeNode
// root 中一定有 node.Val = target
func GetTargetNode(root *TreeNode, target int) *TreeNode {
	if root == nil || root.Val == target {
		return root
	}

	res := GetTargetNode(root.Left, target)
	if res != nil {
		return res
	}
	return GetTargetNode(root.Right, target)
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}

// Tree2Preorder 把 二叉树 转换成 preorder 的切片
func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)

	return res
}

// Equal return ture if tn == a
func (tn *TreeNode) Equal(a *TreeNode) bool {
	if tn == nil && a == nil {
		return true
	}

	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}

	return tn.Left.Equal(a.Left) && tn.Right.Equal(a.Right)
}

// Tree2ints 把 *TreeNode 按照行还原成 []int
func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}
