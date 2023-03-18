import java.util.List;

public class Main {
    public static void main(String[] args) {
        RecoverBST99.TreeNode node2 = new RecoverBST99.TreeNode(2, null, null);
        RecoverBST99.TreeNode node1 = new RecoverBST99.TreeNode(3, null, node2);
        RecoverBST99.TreeNode node0 = new RecoverBST99.TreeNode(1, node1, null);

        RecoverBST99.Solution solution = new RecoverBST99.Solution();
        solution.preorderTraverseTree(node0);
        System.out.println(solution.vals);
    }
}
