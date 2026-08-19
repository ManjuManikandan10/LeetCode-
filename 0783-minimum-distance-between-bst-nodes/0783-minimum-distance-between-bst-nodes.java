/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {

    private int res = Integer.MAX_VALUE;
    // prev tracks the value of the immediately preceding node in sorted order 🎯
    private Integer prev = null;

    public int minDiffInBST(TreeNode root) {
        inorder(root);
        return res;
    }

    // In-order traversal processes nodes in ascending order (sorted sequence) 📈
    private void inorder(TreeNode root) {
        if (root == null) return;

        inorder(root.left);

        // Compare current node with previous adjacent node in sorted order 🔍
        if (prev != null) {
            res = Math.min(res, root.val - prev);
        }
        prev = root.val; // Update prev to current node's value

        inorder(root.right);
    }
}