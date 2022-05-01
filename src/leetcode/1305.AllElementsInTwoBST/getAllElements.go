package main
import "go_practice/structs"

type TreeNode = structs.BTree
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
		res = append(res, nums1...)
	}
	if j<n {
		res = append(res, nums2...)
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

func main() {

}