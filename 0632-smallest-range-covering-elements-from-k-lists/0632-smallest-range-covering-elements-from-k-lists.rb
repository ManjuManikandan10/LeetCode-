# @param {Integer[][]} nums
# @return {Integer[]}
def smallest_range(nums)

    # populate pq
    pq = []
    nums.each_with_index {|n, i| pq.hpush([n.first,i, 0])}

    # first valid range
    ans = [pq.first.first, pq.last.first]

    loop do
        _, i, j = pq.shift

        # if we are going beyond the list boundaries the answer is ready
        return ans if j == nums[i].size - 1
        
        # put list back with new leading element
        pq.hpush([nums[i][j+1], i, j+1])
        
        # only update answer if the new range is less than current range
        if pq.last.first - pq.first.first < ans[1] - ans[0]
            ans = [pq.first.first, pq.last.first]
        end
    end
end

# Use this to substitute PriorityQueue
class Array
    def hpush(x)
        idx = bsearch_index {|y| y[0]>x[0]} || size
        insert(idx,x)
    end
end