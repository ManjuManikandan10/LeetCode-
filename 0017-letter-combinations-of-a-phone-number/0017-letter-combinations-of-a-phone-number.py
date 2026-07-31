class Solution:
    def letterCombinations(self, digits):
        if not digits:
            return []

        # Mapping of digits to letters
        digit_to_char = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz"
        }

        res = []

        def backtrack(index, path):
            # If path is complete, add to result
            if index == len(digits):
                res.append("".join(path))
                return
            
            # Explore letters for current digit
            for char in digit_to_char[digits[index]]:
                path.append(char)
                backtrack(index + 1, path)
                path.pop()  # backtrack

        backtrack(0, [])
        return res
