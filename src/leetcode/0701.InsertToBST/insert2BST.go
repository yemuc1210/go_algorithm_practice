package lt701
import "structs"
type TreeNode = structs.TreeNode
/*
给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 
输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。

*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    p := root
    for p != nil {
        if val < p.Val {
            if p.Left == nil {
                p.Left = &TreeNode{Val: val}
                break
            }
            p = p.Left
        } else {
            if p.Right == nil {
                p.Right = &TreeNode{Val: val}
                break
            }
            p = p.Right
        }
    }
    return root
}