def longest_univalue_path(root)

    return 0 if root.nil?

    @result = 0 

    dfs(root)
    
    return @result 
end 

def dfs(node)
    if node.left.nil? and node.right.nil?
        return 0 
    end 

    lmax = 0 
    rmax = 0 

    if node.left
        if node.left.val == node.val
            lmax = dfs(node.left)+1
        else
            lmax = dfs(node.left)
            lmax = 0            
        end 
    end 

    if node.right
        if node.right.val == node.val
            rmax = dfs(node.right)+1
        else
            rmax = dfs(node.right)
            rmax = 0            
        end 
    end 

    @result = [@result, rmax+lmax].max 

    return [lmax, rmax].max 
end 
