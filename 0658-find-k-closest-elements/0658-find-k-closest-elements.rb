def find_closest_elements(arr, k, x)
    start = (0 ... arr.size - k).bsearch { |i| x - arr[i] <= arr[i + k] - x } || arr.size - k
    arr[start ... start + k]
end