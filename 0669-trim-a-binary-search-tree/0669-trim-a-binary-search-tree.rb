def trim_bst(root, l, r)
    # p "input: #{root ? root.val : 'nil'}"
    root = find_root(root, l, r)
    # p "set: #{root ? root.val : 'nil'}"
    if root
        root.left = trim_bst(root.left, l, r)
        root.right = trim_bst(root.right, l, r)
    end
    root
end
def find_root(root,l,r)
    return root if root.nil?
    return root if root.val.between?(l,r)
    return find_root(root.right,l,r) if root.val < l
    return find_root(root.left,l,r) if r < root.val
end