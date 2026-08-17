class Solution:
    def canReach(self, start: list[int], target: list[int]) -> bool:
        color1: int = (start[0] + start[1]) % 2  
        color2: int = (target[0] + target[1]) % 2
        
        if color1 == color2:
            return True
        else: 
            return False  