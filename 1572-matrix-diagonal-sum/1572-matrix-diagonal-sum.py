class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        count = 0
        n = len(mat)
        for i in range(n):
            count += mat[i][i]
            count += mat[i][n-1-i]
        if n % 2 == 1:
            count -= mat[n//2][n//2]
        return count