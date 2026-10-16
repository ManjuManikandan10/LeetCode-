class Solution:
    def canFormArray(self, arr: List[int], pieces: List[List[int]]) -> bool:
        mp = {x[0]: x for x in pieces}
        i = 0
        while i < len(arr): 
            if (x := arr[i]) not in mp or mp[x] != arr[i:i+len(mp[x])]: return False 
            i += len(mp[x])
        return True 