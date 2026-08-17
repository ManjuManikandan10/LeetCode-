class Solution:
    def maxDigitRange(self, nums: list[int]) -> int:

        digRnge = [0] * 10

        for num in nums:
            mx = mn = num %10

            n = num
            while n:
                n, d = divmod(n, 10)
                if   mx < d: mx = d
                elif mn > d: mn = d

            digRnge[mx - mn]+= num 

        return next(x for x in digRnge[::-1] if x)
      