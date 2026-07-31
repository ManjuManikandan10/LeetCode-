class Solution:
    def isPalindrome(self, x):
        # Negative numbers are not palindromes
        if x < 0 or (x % 10 == 0 and x != 0):
            return False
        
        reverted = 0
        while x > reverted:
            reverted = reverted * 10 + x % 10
            x //= 10
        
        # Check if the number is palindrome
        return x == reverted or x == reverted // 10
  