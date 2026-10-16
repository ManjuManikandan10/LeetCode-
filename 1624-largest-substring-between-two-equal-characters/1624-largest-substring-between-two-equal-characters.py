class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        
        st = {}        
        res = -1

        for i in range(len(s)):

            if s[i] not in st:
                st[s[i]] = i
        
            res = max(res, i - st[s[i]]-1)
        
        return res