class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function findNthDigit($n) {
        $digitLength = 1;
        $count = 9;
        $start = 1;

        while ($n > $digitLength * $count) {
            $n -= $digitLength * $count;
            $digitLength++;
            $count *= 10;
            $start *= 10;
        }

        $start += floor(($n - 1) / $digitLength);
        $numberStr = strval($start);
        
        return intval($numberStr[($n - 1) % $digitLength]);
    }
}