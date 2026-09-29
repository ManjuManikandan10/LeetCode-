class Solution:
    def reverseParentheses(self, s: str) -> str:
        def funct(l, r):
            a = ""
            i = l
            while l <= i <= r:
                if s[i].isalpha():
                    a += s[i]
                    i += 1
                elif s[i] == "(":
                    c = 1
                    j = i + 1
                    while c != 0:
                        if s[j] == ")":
                            c -= 1
                        elif s[j] == "(":
                            c += 1
                        j += 1
                    b = funct(i + 1, j - 2)
                    i = j
                    a = a + b
            return a[::-1]
        return funct(0, len(s) - 1)[::-1]