class Solution:
    def minSubarray(self, nums, p):
        total = sum(nums)
        target = total % p

        # Already divisible
        if target == 0:
            return 0

        # remainder -> latest index
        prefix = {0: -1}
        current = 0
        answer = len(nums)

        for i, num in enumerate(nums):
            current = (current + num) % p

            # We need a subarray with remainder = target
            needed = (current - target) % p

            if needed in prefix:
                answer = min(answer, i - prefix[needed])

            # Store latest index
            prefix[current] = i

        # Cannot remove the entire array
        if answer == len(nums):
            return -1

        return answer