def insert_into_bst(root, val)
    new_node = TreeNode.new(val)

    return new_node if root.nil?

    current = root

    while current.nil? == false
        return root if current.val === val

        if val < current.val
            if current.left.nil?
                current.left = new_node
                break
            end

            current = current.left
        else
            if current.right.nil?
                current.right = new_node
                break
            end

            current = current.right
        end
    end

    return root
end