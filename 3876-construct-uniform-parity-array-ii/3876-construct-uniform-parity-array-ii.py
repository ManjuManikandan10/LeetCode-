class Solution:
    def uniformArray(self, nums1: list[int]) -> bool:
        parity = min(nums1) % 2
        min_odd = min((i for i in nums1 if i % 2 != 0), default = 10**9 + 1)
        for x in nums1:
            if x % 2 != parity and min_odd >= x:
                return False
        return True