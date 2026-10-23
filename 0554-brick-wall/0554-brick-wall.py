class Solution:
    def leastBricks(self, wall: List[List[int]]) -> int:
        breaks = defaultdict(int)

        for row in wall:
            position = 0
            for brick in row[:-1]: 
                position += brick
                breaks[position] += 1

        max_breaks = max(breaks.values(),default=0)
        res = len(wall) - max_breaks

        return res