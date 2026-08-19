class Solution {

    public List<Integer> postorder(Node root) {

        List<Integer> arr = new ArrayList<>();

        if (root == null) {
            return arr;
        }

        for (Node child : root.children) {
            arr.addAll(postorder(child));
        }

        arr.add(root.val);

        return arr;
    }
}