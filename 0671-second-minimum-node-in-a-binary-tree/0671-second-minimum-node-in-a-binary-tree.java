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
    List<Integer>list=new ArrayList<>();
    public int findSecondMinimumValue(TreeNode root) {
        helper(root);
        Collections.sort(list);
        int last=list.get(0);
        for(int val:list){
            if(val!=last)return val;
        }
        return -1;
    }
    public void helper(TreeNode root){
        list.add(root.val);
        if(root.left!=null)helper(root.left);
        if(root.right!=null)helper(root.right);
    }
}