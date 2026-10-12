class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        m=len(grid)
        n=len(grid[0])
        ans=[]
        count=0
        for i in range(m):
            for j in range(n):
                ans.append(grid[i][j])
        for i in ans:
            if i <0:
                count+=1       
        return count
        
        