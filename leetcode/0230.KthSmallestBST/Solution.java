import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Stack;

class Solution {
    // 230. 二叉搜索树中第K小的元素 Java
    public int kthSmallest(TreeNode root, int k) {
        // 得到中序遍历
        ArrayList<Integer> list = new ArrayList<>();
        Stack<TreeNode> sk = new Stack<>();
        while(root!=null || !sk.isEmpty()){
            while(root !=null) {
                sk.push(root);
                root = root.left;
            }
            // 中
            root = sk.pop();
            list.add(root.val);
            // 右
            root = root.right;
        }
        // 返回
        return list.get(k-1);
    }
}

public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(){}
    TreeNode(int val) {
        this.val = val;
    }
    TreeNode(int val, TreeNode left,TreeNode right){
        this.val = val;
        this.right = right;
        this.left = left;
    }
}