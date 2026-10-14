class Solution:
    def threeConsecutiveOdds(self, arr: List[int]) -> bool:

        for a, b, c in zip(arr, arr[1:], arr[2:]):
            if a & b & c & 1: return True
        return False