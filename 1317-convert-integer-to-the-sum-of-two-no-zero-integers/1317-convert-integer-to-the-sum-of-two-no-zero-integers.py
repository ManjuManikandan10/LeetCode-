class Solution:
    def getNoZeroIntegers(self, n: int) -> List[int]:
        a=n//2
        b=n-a
        """
        we dont want any 0 number in both of these!!
        """
        # return [a,b]
        # based on this we see just O(N) for the constriant beign N
        for i in range(1,n):
            if '0' not in str(i) and '0' not in str(n-i):
                return [i,n-i]
        # since testcases are such that atleast one valid exist no need to worry!1