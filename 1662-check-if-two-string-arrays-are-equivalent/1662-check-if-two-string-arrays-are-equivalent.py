class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        length1 = len(word1)
        length2 = len(word2)
        for i in range(1,length1):
            word1[0] = word1[0] + word1[i]

        for i in range(1, length2):
            word2[0] = word2[0] + word2[i]

        if(word1[0] == word2[0]):
            return True
        else:
            return False