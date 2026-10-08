class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        count = 0
        res = 0
        for char in s:
            if char == "(":
                count += 1
            else:
                count -= 1
            if count < 0:
                res += 1
                count = 0
        if count > 0:
            res += count
        return abs(res)