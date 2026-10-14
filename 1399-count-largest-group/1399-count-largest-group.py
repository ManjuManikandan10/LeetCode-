class Solution:
    def countLargestGroup(self, n):
        groups = {}

        for num in range(1, n + 1):
            digit_sum = sum(int(d) for d in str(num))
            groups[digit_sum] = groups.get(digit_sum, 0) + 1

        max_size = max(groups.values())

        return sum(1 for size in groups.values() if size == max_size)