class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        left = 0 
        answer = 0
        max_freq = 0
        freq = defaultdict(int) 

        for right in range(len(s)):
            freq[s[right]] += 1

            max_freq = max(max_freq, freq[s[right]])

            while (right - left + 1) - max_freq > k:
                freq[s[left]] -= 1
                left += 1

            answer = max(answer, right - left + 1)
        return answer        
