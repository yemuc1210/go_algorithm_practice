import java.util.HashSet;

class TreeNode {
    int val;
    TreeNode left;
      TreeNode right;
      TreeNode() {}
      TreeNode(int val) { this.val = val; }
      TreeNode(int val, TreeNode left, TreeNode right) {
          this.val = val;
          this.left = left;
          this.right = right;
      }
  }
 
class Solution {
    public boolean findTarget(TreeNode root, int k) {
        HashSet<Integer> hashSet = new HashSet<Integer>();
        return preOrder(root,hashSet,k);
   }
   public boolean preOrder(TreeNode root, HashSet<Integer> hashSet, int target) {
       if(root == null ){
           return false;
       }
       if(hashSet.contains(target - root.val)) {
           return true;
       }
       hashSet.add(root.val);
       return preOrder(root.left, hashSet, target) || preOrder(root.right, hashSet, target);
   }
}
