def find_duplicate(paths)
    contents = Hash.new { |h,k| h[k] = Array.new }
    paths.each do |group| 
        arr = group.split" "
        path = arr.shift
        arr.each do |file|
            par_idx = file.index('(')
            contents[file[par_idx + 1...-1]] << ( path + "/" + file[...par_idx] )
        end
    end
    contents.values.select { _1.length > 1}
end