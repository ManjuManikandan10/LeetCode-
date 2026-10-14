class Solution:
    def minStartValue(self, nums: List[int]) -> int:
        min_prefix =0
        count =0
        for num in nums:
            count += num
            min_prefix=min(min_prefix,count)
        return 1-min_prefix
        