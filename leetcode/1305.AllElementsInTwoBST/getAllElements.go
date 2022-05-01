package main

import (
	"fmt"
	"go_practice/structs"
)

type TreeNode = structs.TreeNode
// type TreeNode struct{
// 	Val int
// 	Left *TreeNode
// 	Right *TreeNode
// }

// 两个BST 返回两个数中所有的数，并升序排
// 中序遍历，获得一个树得有序序列，然后merge两个有序序列
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var nums1 = []int{}
	var nums2 = []int{}
	inorder(root1,&nums1)
	inorder(root2,&nums2)
	return merge(nums1,nums2)
}
func merge(nums1,nums2 []int) []int{
	var res []int
	m,n := len(nums1),len(nums2)
	var i,j int
	for i<m && j<n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		}else{
			res = append(res, nums2[j])
			j++
		}
	}
	if i<m {
		for i<m {
			res = append(res, nums1[i])
			i++
		}
	}
	if j<n {
		for j<n {
			res = append(res, nums2[j])
			j++
		}
	}

	return res
}
func inorder(root *TreeNode, res *[]int){
	if root == nil {
		return 
	}
	inorder(root.Left,res)
	*res = append(*res, root.Val)
	inorder(root.Right,res)
}
// 中序遍历的非递归
func inorder1(root *TreeNode) {
	// 中序遍历的顺序：左中右
	if root == nil {
		return
	}
	var res []int 
	sk := []*TreeNode{root}
	for len(sk)>0 {
		curNode := sk[len(sk)-1]
		// 访问左
		for curNode.Left!=nil {
			curNode = curNode.Left
			sk = append(sk, curNode)
		}
		// 中
		res = append(res, curNode.Val)
		// 右
		sk = append(sk, curNode.Right)
	}
	fmt.Println(res)
}

func main() {
	root1 := structs.Ints2TreeNode([]int{2,1,4})
	root2 := structs.Ints2TreeNode([]int{1,0,3})
	res := getAllElements(root1,root2)
	fmt.Println(res)

	root := structs.Ints2TreeNode([]int{2,1,4})
	if root == nil {
		return
	}
	res = []int{}
	sk := []*TreeNode{root}
	for len(sk)>0 {
		curNode := sk[len(sk)-1]
		// 访问左
		for curNode.Left!=nil {
			curNode = curNode.Left
			sk = append(sk, curNode)
		}
		// 中
		res = append(res, curNode.Val)
		// 右
		sk = append(sk, curNode.Right)
	}
	fmt.Println(res)
}