class Solution:
    def busyStudent(self, startTime: List[int], 
                          endTime: List[int], queryTime: int) -> int:

        inInterval = lambda x: x[0] <= queryTime <= x[1]
 
        return sum(map(inInterval, zip(startTime, endTime)))