class Solution:
    def reverse(self, x):
        INT_MAX = 2**31 - 1
        INT_MIN = -2**31
        
        result = 0
        negative = x < 0
        x = abs(x)
        
        while x != 0:
            pop = x % 10
            x //= 10
            
            # Check for overflow before actually appending
            if result > (INT_MAX - pop) // 10:
                return 0
            
            result = result * 10 + pop
        
        return -result if negative else result
