class Solution {

    /**
     * @param String $n
     * @return String
     */
    function nearestPalindromic($n) {
        $num = (int)$n;
        $numLength = strlen($n);

        // Generate boundary palindromes
        $boundaryPalindromes = [
            (10 ** $numLength + 1) . '',
            (10 ** ($numLength - 1) - 1) . '',
        ];

        // Generate palindromes around the input number
        $prefix = (int)substr($n, 0, floor(($numLength + 1) / 2));
        foreach ([ $prefix - 1, $prefix, $prefix + 1 ] as $start) {
            $firstHalf = (string)$start;
            $palindrome = $this->_makePalindrome($firstHalf, $numLength);
            $boundaryPalindromes[] = $palindrome;
        }

        // Remove the input number from the array if it's present
        $boundaryPalindromes = array_diff($boundaryPalindromes, [$n]);

        // Find the nearest palindrome
        $nearest = null;
        $minDiff = INF;
        foreach ($boundaryPalindromes as $palindrome) {
            $palindromeInt = (int)$palindrome;
            $diff = abs($palindromeInt - $num);
            if ($diff < $minDiff || ($diff === $minDiff && $palindromeInt < (int)$nearest)) {
                $nearest = $palindrome;
                $minDiff = $diff;
            }
        }

        return $nearest;
    }

    private function _makePalindrome($firstHalf, $length) {
        if ($length % 2) {
            return $firstHalf . strrev(substr($firstHalf, 0, -1));
        } else {
            return $firstHalf . strrev($firstHalf);
        }
    }
}