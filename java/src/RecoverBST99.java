import java.util.List;
import java.util.ArrayList;

public class RecoverBST99 {
    public static class TreeNode {
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
    public static class Solution {
        List<Integer> vals = new ArrayList();
        public void recoverTree(TreeNode root) {

        }
        void preorderTraverseTree(TreeNode node) {
            if (node == null) {
                return;
            }
            preorderTraverseTree(node.left);
            vals.add(node.val);
            preorderTraverseTree(node.right);
        }
    }
}
