class Solution:
    def rearrangeString(self, s: str, x: str, y: str) -> str:

        return ''.join([y * s.count(y), s.replace(y,'')])

        