class Solution:
    def isValid(self, code: str) -> bool:
        stack = []
        i = 0
        while i < len(code):
            if i > 0 and not stack:
                return False
            if code[i:i+9] == "<![CDATA[":
                i = code.find("]]>", i + 9)
                if i == -1:
                    return False
                i = i + 3
            elif code[i: i + 2] == "</" and stack:
                endTag = stack.pop() + ">"
                if code[i + 2: i + 2 + len(endTag)] != endTag:
                    return False
                i = i + len(endTag) + 2
            elif code[i] == "<":
                x = code.find(">", i + 2, i + 11)
                if x == -1 or not (code[i + 1: x].isalpha() and code[i + 1: x].isupper()):
                    return False
                stack.append(code[i + 1: x])
                i = x + 1
            else:
                i += 1
        return not stack