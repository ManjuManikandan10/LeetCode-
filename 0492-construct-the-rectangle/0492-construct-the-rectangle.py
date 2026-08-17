from math import sqrt

class Solution:
    def constructRectangle(self, area: int) -> list[int]:
        w = int(sqrt(area))
        while area%w:
            w -= 1
        return [area//w, w]