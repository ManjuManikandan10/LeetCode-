class Solution:
    def firstStableIndex(self, nums: list[int], k: int) -> int:
        minArr = list()
        freq = dict()
        for num in nums:
            minArr.append(num)
            freq[num] = freq.get(num, 0) + 1
        minArr.sort()
        maxEle, minInd = -1, 0
        for i in range(0, len(nums)):
            maxEle = max(maxEle, nums[i])
            if maxEle - minArr[minInd] <= k:
                return i
            freq[nums[i]] -= 1
            while minInd < len(nums) and freq[minArr[minInd]] <= 0:
                minInd+=1
            
        return -1