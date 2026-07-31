class Solution:
    def myAtoi(self, s):
        INT_MAX = 2**31 - 1
        INT_MIN = -2**31
        
        i = 0
        n = len(s)
        # Skip leading whitespaces
        while i < n and s[i] == ' ':
            i += 1
        
        # Check if empty string after removing spaces
        if i == n:
            return 0
        
        # Check sign
        sign = 1
        if s[i] == '+':
            i += 1
        elif s[i] == '-':
            sign = -1
            i += 1
        
        result = 0
        while i < n and s[i].isdigit():
            digit = int(s[i])
            
            # Check for overflow
            if result > (INT_MAX - digit) // 10:
                return INT_MAX if sign == 1 else INT_MIN
            
            result = result * 10 + digit
            i += 1
        
        return sign * result
