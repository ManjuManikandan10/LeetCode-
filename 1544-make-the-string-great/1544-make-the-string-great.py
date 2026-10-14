class Solution:
    def makeGood(self, s: str) -> str:
        # S1: Brute Force
        # i = 0
        # while i < len(s) - 1:
        #     if s[i].lower() == s[i+1].lower() and s[i] != s[i+1]:
        #         s = s[:i] + s[i+2:]
        #         i = 0
        #     else:
        #         i += 1
        # return s
        #
        # S2: Stack
        # Time: O(n)
        # Auxiliary space: O(n)
        stack = []
        for char in s:
            if stack and char.lower() == stack[-1].lower() and char != stack[-1]:
                    stack.pop()
            else:
                stack.append(char)
        return "".join(stack)

        