class Solution:
    def longestCommonPrefix(self, strs):
        if not strs:
            return ""
        
        # Start with the first string as the prefix
        prefix = strs[0]
        
        for s in strs[1:]:
            # Shorten the prefix until it matches the start of s
            while not s.startswith(prefix):
                prefix = prefix[:-1]
                if not prefix:
                    return ""
        
        return prefix
