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
     * @param Integer $key
     * @return TreeNode
     */
    function deleteNode($root, $key) {
        if ($root === null) {
            return null; // Key not found, return the tree unchanged
        }

        if ($key < $root->val) {
            // Key is in the left subtree
            $root->left = $this->deleteNode($root->left, $key);
        } elseif ($key > $root->val) {
            // Key is in the right subtree
            $root->right = $this->deleteNode($root->right, $key);
        } else {
            // Node to delete is found
            if ($root->left === null) {
                // Case 1: No left child
                return $root->right;
            } elseif ($root->right === null) {
                // Case 2: No right child
                return $root->left;
            } else {
                // Case 3: Two children
                // Find the inorder successor (smallest value in the right subtree)
                $successor = $this->findMin($root->right);
                $root->val = $successor->val; // Replace current node value with successor value
                $root->right = $this->deleteNode($root->right, $successor->val); // Delete the successor
            }
        }

        return $root;
    }

    private function findMin($node) {
        while ($node->left !== null) {
            $node = $node->left;
        }
        return $node;
    }
}