public class MinimimDepthBT111 {
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
    static class Solution {
        public int minDepth(TreeNode root) {
            if (root == null) {
                return 0;
            }
            int minDepthLeft = minDepth(root.left);
            int minDepthRight = minDepth(root.right);
            int minDepthValid;
            if (minDepthLeft == 0 && minDepthRight != 0) {
                minDepthValid = minDepthRight;
            } else if (minDepthLeft != 0 && minDepthRight == 0) {
                minDepthValid = minDepthLeft;
            } else {
                minDepthValid = (minDepthLeft < minDepthRight) ? minDepthLeft : minDepthRight;
            }
            return minDepthValid + 1;
        }
    }
}
