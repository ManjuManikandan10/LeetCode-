class Solution:
    def containsPattern(self, nums, m, k) -> bool:
        len_nums = len(nums)
        nums_str = ','.join([str(n) for n in nums])
        i = 0
        while (i + m < len_nums):
            pattern = ','.join([','.join([str(n) for n in nums[i:i+m]])] * k)
            if (nums_str.find(pattern) >= 0):
                return True
            i += 1

        return False