class Solution:
    def freqAlphabets(self, s: str) -> str:
        i = 0
        word = ''
        while i < len(s):
            if i + 2 < len(s) and s[i + 2] == '#':
                word += chr(int(s[i:i+2]) + 96)
                i += 3
            else:
                word += chr(int(s[i]) + 96)
                i += 1
        return word