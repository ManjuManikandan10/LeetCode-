class Solution:
    def maxArea(self, height):
        left, right = 0, len(height) - 1
        max_water = 0

        while left < right:
            # Width of the container
            width = right - left
            # Height is limited by the shorter line
            h = min(height[left], height[right])
            max_water = max(max_water, width * h)

            # Move the pointer pointing to the shorter line
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1

        return max_water
