def num_subarray_product_less_than_k(nums, k)
    return 0 if k.zero?
    i, prod = -1, 1
    nums.each_with_index.inject(0) {|res, (v, j)|
        prod *= v
        prod /= nums[i += 1] while i < j && prod >= k
        res + j - i
    }
end