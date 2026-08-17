class Solution:
    def checkGoodInteger(self, n: int) -> bool:
        return sum([int(x)**2 - int(x) for x in str(n)]) >= 50
        