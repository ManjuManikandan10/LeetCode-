class Solution:
    def reorderSpaces(self, text: str) -> str:
        str_len = len(text)
        chars = text.split(" ")
        ch = []
        chlen = 0
        for i in chars:
            if not i:
                continue
            ch.append(i)
            chlen += len(i)
        if len(ch) == 1:
            return ch[0] + (" " * (str_len - len(ch[0]))) 
        
        spaces = (str_len - chlen) // (len(ch)-1)
        at_the_end = (str_len - chlen) % (len(ch)-1)
        res = ''
        for i in range(len(ch)-1):
            res += ch[i] + (" " * spaces)
        res += ch[-1] + (" "*at_the_end)
        return res
