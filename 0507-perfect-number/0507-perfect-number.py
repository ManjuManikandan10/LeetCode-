class Solution:
    def checkPerfectNumber(self, num: int) -> bool:
        #If the given number is 1, return False since 1 cannot be a perfect number.
        if num == 1:
            return False

        #Initialize a variable "ans" as 1 to store the running sum.
        ans = 1

        #Iterate through all numbers from 2 up to the square root of the given number.
        for i in range(2, int(num**0.5) + 1):
            #Check if the current number is a divisor of the given number.
            if num % i == 0:
                #If it is a divisor, add both the divisor and the quotient to the running sum.
                ans += i + num // i
        #After the loop, compare the running sum with the given number. 
        #If they are equal, return True (the number is a perfect number); otherwise, return False.
        return ans == num