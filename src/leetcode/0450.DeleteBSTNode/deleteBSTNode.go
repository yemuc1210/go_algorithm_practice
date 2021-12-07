package lt450
import "structs"
type TreeNode = structs.TreeNode
// func deleteNode(root *TreeNode, key int) *TreeNode {
// 	if root == nil{
// 		return nil
// 	}
// 	if key < root.Val {
// 		//左子树
// 		root.Left = deleteNode(root.Left, key)
// 	}else if key > root.Val{
// 		root.Right = deleteNode(root.Right, key)
// 	}else{
// 		//root待删除
// 		if root.Left == nil{
// 			return root.Right
// 		}else if root.Right == nil {
// 			return root.Left
// 		}else{
// 			//左右子树都存在
// 			successor := min(root.Right)  //选择右子树中最左叶子节点
// 			successor.Right = deleteMin(root.Right)
// 			successor.Left = root.Left
// 			return successor
// 		}
// 	}
// 	return nil
// }
// func min(node *TreeNode)*TreeNode{
// 	if node.Left == nil {
// 		return node
// 	}
// 	return min(node.Left)
// }
// func deleteMin(node *TreeNode)*TreeNode{
// 	if node.Left == nil{
// 		return node.Right
// 	}
// 	node.Left = deleteMin(node.Left)
// 	return node
// }
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key > root.Val {
        root.Right = deleteNode(root.Right, key)
        return root
    } else if key < root.Val {
        root.Left = deleteNode(root.Left, key)
        return root
    } 
    if root.Left == nil && root.Right == nil {
        return nil
    }
    if root.Left == nil {
        return root.Right
    } else if root.Right == nil {
        return root.Left
    }
    cur := root.Right
    for cur.Left != nil {
        cur = cur.Left
    }
    root.Val = cur.Val
    root.Right = deleteNode(root.Right, cur.Val)
    return root
}
