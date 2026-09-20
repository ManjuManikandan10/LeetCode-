class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function reverseDegree($s) {
        $sum = 0;
        for ($i = 0; $i < strlen($s); $i++) {
            $sum += (26 - ord($s[$i]) + 97) * ($i+1);
        }
        return $sum;
    }
}