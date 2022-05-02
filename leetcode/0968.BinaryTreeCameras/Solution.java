class Solution {
    // 968. 监控二叉树 Java
    // 给定一个二叉树，我们在树的节点上安装摄像头。
    // 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
    // 计算监控树的所有节点所需的最小摄像头数量。
    int cameraCnt = 0;
    public int minCameraCover(TreeNode root) {
        // 节点：根 叶子 监控
        if(root==null) {
            return 0;
        }
        if(dfs(root)==2){
            cameraCnt++;
        }
        return cameraCnt;
    }
    // 0:安装  1：节点可观测，但没有安装  2：节点不可观测
    int dfs(TreeNode node) {
        if(node==null){
            return 1;
        }
        int left = dfs(node.left),right=dfs(node.right);
        if(left==2 || right==2){
            cameraCnt++;
            return 0;
        }else if(left==0 || right==0){
            return 1;
        }else{
            return 2;
        }
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