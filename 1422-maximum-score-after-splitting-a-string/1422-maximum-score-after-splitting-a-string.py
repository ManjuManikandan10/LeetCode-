class Solution:
    def maxScore(self, s: str) -> int:
        total_zeroes = s.count('0')
        total_left_zeroes = 0
        ans = -1
        for i in range(1, len(s)):
            if s[i-1] == '0':
                total_left_zeroes += 1
            total_right_zeroes = total_zeroes - total_left_zeroes 
            ans = max(ans, total_left_zeroes + (len(s) - i - total_right_zeroes))
        return ans