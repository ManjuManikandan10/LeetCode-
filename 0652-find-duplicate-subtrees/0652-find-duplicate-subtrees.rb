def find_duplicate_subtrees(root)
    @path = {}
    @duplicates = []
    
    find_duplicates(root)
    
    @duplicates
end

def find_duplicates(root)
    return "" if root.nil?
    
    path = [root.val,find_duplicates(root.left),find_duplicates(root.right)].join(',')
    
    if @path[path]
        @duplicates << root if @path[path] == 1
        @path[path] += 1
    else
        @path[path] = 1
    end
    
    path
end