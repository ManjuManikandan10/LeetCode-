
class Solution {
    public List<Integer> preorder(Node root) {
         List<Integer> sta = new ArrayList<>();
        dfs(root,sta);
        return sta;
    }
    public void dfs(Node root, List<Integer> sta){
        if(root==null) return;
        sta.add(root.val);
        for(Node child:root.children){
        dfs(child,sta);
    }
}}
    