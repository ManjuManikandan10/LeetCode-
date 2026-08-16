class Solution:
    def waysToMakeFair(self, nums):
        total_even = 0
        total_odd = 0

        # Calculate total even and odd index sums
        for i, num in enumerate(nums):
            if i % 2 == 0:
                total_even += num
            else:
                total_odd += num

        left_even = 0
        left_odd = 0
        answer = 0

        for i, num in enumerate(nums):

            # Remove current element from the right side
            if i % 2 == 0:
                total_even -= num
            else:
                total_odd -= num

            # After removing nums[i], elements on the right
            # change their index parity.
            if left_even + total_odd == left_odd + total_even:
                answer += 1

            # Add current element to the left side
            if i % 2 == 0:
                left_even += num
            else:
                left_odd += num

        return answer