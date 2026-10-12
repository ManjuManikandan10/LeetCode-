class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        ans=0
        for i in range(len(arr1)):
            count=0
            for j in range(len(arr2)):
                if abs(arr1[i]-arr2[j]) <= d:
                    count+=1
            if count==0:
                ans+=1
        return ans