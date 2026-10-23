class Solution:
    def smallestGoodBase(self, n: str) -> str:
        #find the smallest base(k) of n where all are 1 in k 
        #we dont know how many ones we have 


        #a number can have so many ones lets assumes m ones  

        #1*(k^0)+1*(k^1)+... 1*(k^(m-1))
        #this is summation of a GP 
        #1+k+k^2+...+k^m 
        #(k^(m)-1)/(k-1)

        #find the base k for the n 
        n=int(n)

        #(k**(m)-1)//(k-1)-n==0
        for m in range(1,101):
            low=2
            high=10**18 
            while low<=high:
                mid=(low+high)>>1 
                val=(mid**(m)-1)//(mid-1)-n
                if val==0:
                    ans=mid 
                    high=mid-1
                if val>0:
                    high=mid-1 
                else:
                    low=mid+1
        return str(ans)