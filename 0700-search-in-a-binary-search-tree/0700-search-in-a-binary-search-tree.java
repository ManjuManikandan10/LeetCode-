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
    public TreeNode searchBST(TreeNode root, int val) {
        // Base Case: Target found or node doesn't exist 🎯
        if (root == null) return root;
        
        // Target is smaller -> search left subtree ⬅️
        if (val < root.val) {
            return searchBST(root.left, val);
        } 
        // Target is larger -> search right subtree ➡️
        else if (val > root.val) {
            return searchBST(root.right, val);
        }
        
        // Target node matched! Return subtree rooted at target 🌳
        return root;
    }
}