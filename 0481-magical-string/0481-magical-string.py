class Solution:
    def magicalString(self, n):
        if n == 1:
            return 1

        s = [1, 2, 2]
        i = 2
        num = 1
        count = 1

        while len(s) < n:
            # Add num according to the current group length
            for _ in range(s[i]):
                s.append(num)

                if num == 1 and len(s) <= n:
                    count += 1

            # Switch between 1 and 2
            num = 3 - num
            i += 1

        return count