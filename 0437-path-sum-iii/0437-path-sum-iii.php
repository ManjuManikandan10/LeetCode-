/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @param Integer $targetSum
     * @return Integer
     */
    function pathSum($root, $targetSum) {
        $prefixSumCount = [0 => 1]; // To handle cases where the path starts from the root
        return $this->dfs($root, $targetSum, 0, $prefixSumCount);
    }

    private function dfs($node, $targetSum, $currentSum, &$prefixSumCount) {
        if ($node === null) {
            return 0;
        }

        // Update the current prefix sum
        $currentSum += $node->val;

        // Count paths ending at the current node that sum to targetSum
        $paths = $prefixSumCount[$currentSum - $targetSum] ?? 0;

        // Update the prefix sum count
        $prefixSumCount[$currentSum] = ($prefixSumCount[$currentSum] ?? 0) + 1;

        // Recurse on left and right children
        $paths += $this->dfs($node->left, $targetSum, $currentSum, $prefixSumCount);
        $paths += $this->dfs($node->right, $targetSum, $currentSum, $prefixSumCount);

        // Backtrack: Remove the current prefix sum from the hash map
        $prefixSumCount[$currentSum]--;

        return $paths;
    }
}