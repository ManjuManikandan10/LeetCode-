class Solution:
    def findNumbers(self, nums: list[int]) -> int:
        nums=[str(val) for val in nums]
        res=0
        for val in nums:
            if len(val)%2==0:
                res+=1
        return res