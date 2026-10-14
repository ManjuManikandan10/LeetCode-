class Solution:
    def canBeEqual(self, target: List[int], arr: List[int]) -> bool:
        target = Counter(target)
        arr = Counter(arr)
        for x in arr:
            if x not in target or target[x]!=arr[x]:
                return False
        return True