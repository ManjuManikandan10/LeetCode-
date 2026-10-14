class Solution:
    def reformat(self, s: str) -> str:
        nums = [c for c in s if c.isnumeric()]
        alph = [c for c in s if c.isalpha()]
                
        if abs(len(nums) - len(alph)) > 1:
            return ''
        
        a, b = (nums, alph) if len(nums) <= len(alph) else (alph, nums)
        return ''.join(c for pair in itertools.zip_longest(b, a) for c in pair if c)