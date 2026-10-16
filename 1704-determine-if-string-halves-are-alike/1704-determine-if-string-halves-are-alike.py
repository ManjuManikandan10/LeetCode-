class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        alike = {'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
        count_a = 0
        count_b = 0

        a = s[:len(s)//2]
        b = s[len(s)//2:]

        for a_like in a:
            if a_like in alike:
                count_a += 1

        for b_like in b:
            if b_like in alike:
                count_b += 1

        return count_a == count_b               

       