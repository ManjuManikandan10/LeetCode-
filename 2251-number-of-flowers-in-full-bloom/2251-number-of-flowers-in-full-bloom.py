class Solution:
    def fullBloomFlowers(self, flowers: List[List[int]],
                               people: List[int]) -> List[int]:

        (beg, end), ans  = map(sorted, zip(*flowers)), []

        for person in people:
            ans.append(bisect_right(beg, person) - 
                       bisect_left (end, person))
        return ans