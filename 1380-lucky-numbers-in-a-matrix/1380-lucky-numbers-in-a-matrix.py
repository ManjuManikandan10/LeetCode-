class Solution:
    def luckyNumbers(self, matrix: List[List[int]]) -> List[int]:
        m = []
        r = []
        for i in range(len(matrix)):
            m.append(min(matrix[i]))
        for j in range(len(matrix[0])):
            c = matrix[0][j]
            for i in range(len(matrix)):
                c = max(c, matrix[i][j])
            r.append(c)
        a = []
        for x in m:
            if x in r:
                a.append(x)
        return a