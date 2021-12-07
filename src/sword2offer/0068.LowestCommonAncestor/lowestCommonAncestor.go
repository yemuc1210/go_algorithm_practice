package offer68
import "structs"
type TreeNode = structs.TreeNode
// class Solution {
// 	public:
// 		TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
// 			if(p == NULL || q == NULL || root == NULL){
// 				return NULL;
// 			}
// 			if(p->val < root->val && q->val < root->val ) {
// 				return lowestCommonAncestor(root->left, p ,q);
// 			}
// 			if(p->val > root->val && q->val > root->val ) {
// 				return lowestCommonAncestor(root->right, p ,q);
// 			}
// 			return root;
// 		}
// 	};

/*
当我们分别得到了从根节点到 pp 和 qq 的路径之后，我们就可以很方便地找到它们的最近公共祖先了。
显然，pp 和 qq 的最近公共祖先就是从根节点到它们路径上的「分岔点」，也就是最后一个相同的节点。


*/

func getPath(root, target *TreeNode) (path []*TreeNode) {
    node := root
    for node != target {
        path = append(path, node)
        if target.Val < node.Val {
            node = node.Left
        } else {
            node = node.Right
        }
    }
    path = append(path, node)
    return
}

func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
    pathP := getPath(root, p)
    pathQ := getPath(root, q)
    for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
        ancestor = pathP[i]
    }
    return
}
