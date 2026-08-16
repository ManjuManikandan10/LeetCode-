class Solution:
    def minimumAbsDifference(self, arr):
        arr.sort()
        
        min_diff = float('inf')
        result = []
        
        # First pass: find minimum difference
        for i in range(1, len(arr)):
            diff = arr[i] - arr[i-1]
            if diff < min_diff:
                min_diff = diff
        
        # Second pass: collect all pairs with that difference
        for i in range(1, len(arr)):
            if arr[i] - arr[i-1] == min_diff:
                result.append([arr[i-1], arr[i]])
        
        return result
