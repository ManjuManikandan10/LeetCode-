class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        m = -1
        for i in range(len(arr)-1,-1,-1):
            if arr[i]>m :
                m, arr[i] = arr[i], m
            else:
                arr[i] = m
        return arr
            
            
# 1234
# 4321
# 1111
# 14,33,2,3,6