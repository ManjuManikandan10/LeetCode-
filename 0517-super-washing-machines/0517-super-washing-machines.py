class Solution:
    def findMinMoves(self, machines: List[int]) -> int:

        ave, rem = divmod(sum(machines),len(machines))    # <– 1.
        if rem: return -1                                 #
        
        machines = [m - ave for m in machines]            # <– 2.

        return max(max(machines),                         # <– 3.
                   max(map(abs,(accumulate(machines)))))  # <– 4.