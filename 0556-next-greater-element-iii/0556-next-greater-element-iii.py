class Solution:
    def nextGreaterElement(self, n: int) -> int:
        n = str(n)

        # Find the pivot
        i = len(n) - 2

        while i >= 0 and n[i] >= n[i + 1]:
            i -= 1

        # No greater permutation exists
        if i < 0:
            return -1

        # Find the smallest digit greater than n[i]
        j = len(n) - 1

        while n[j] <= n[i]:
            j -= 1

        # Swap
        nums = list(n)
        nums[i], nums[j] = nums[j], nums[i]

        # Make suffix smallest
        nums[i + 1:] = sorted(nums[i + 1:])

        result = int(''.join(nums))

        # 32-bit integer check
        if result > 2147483647:
            return -1

        return result