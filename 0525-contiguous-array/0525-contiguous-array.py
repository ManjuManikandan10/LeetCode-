class Solution:
    def findMaxLength(self, nums: List[int]) -> int:

        acc = list(accumulate(nums, lambda x,y: 
                                     x+2*y-1, initial = 0)) # <-- 1)


        first, latest = defaultdict(int), defaultdict(int)  # 
                                                            #
        for i,n in enumerate(acc):                          # <-- 2)
            if n in first: latest[n] = i                    #
            else: first[n] = i                              #

        return max((latest[n] - first[n] 
                            for n in latest), default = 0)  # <-- 3)
        