class Solution:
    def isValid(self, s):
        stack = []
        # Mapping of closing to opening brackets
        bracket_map = {')': '(', '}': '{', ']': '['}
        
        for char in s:
            if char in bracket_map:
                # Pop from stack if not empty, else dummy value
                top_element = stack.pop() if stack else '#'
                if bracket_map[char] != top_element:
                    return False
            else:
                # Opening bracket, push to stack
                stack.append(char)
        
        # Valid if stack is empty at the end
        return not stack
