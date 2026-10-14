class Solution:
    def maxPower(self, s: str) -> int:
        max=1
        n=0
        m=1
        c=1
        while(m<len(s)):
            if(s[n]==s[m]):
                c+=1
                m+=1
            else:
                if(max<c):
                    max=c
                n=m
                c=0
        if(max<c):
            max=c
        return max