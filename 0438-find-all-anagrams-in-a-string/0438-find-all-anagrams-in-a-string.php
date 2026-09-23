class Solution {
    /**
     * @param String $s
     * @param String $p
     * @return Integer[]
     */
    function findAnagrams($s, $p) {
        $ans = [];
        $n = strlen($s);
        $m = strlen($p);
        if ($n < $m) {
            return $ans;
        }
        $p_count = array_fill(0, 26, 0);
        $s_count = array_fill(0, 26, 0);
        for ($i = 0; $i < $m; $i++) {
            $p_count[ord($p[$i]) - ord('a')]++;
            $s_count[ord($s[$i]) - ord('a')]++;
        }
        if ($p_count == $s_count) {
            $ans[] = 0;
        }
        for ($i = $m; $i < $n; $i++) {
            $s_count[ord($s[$i]) - ord('a')]++;
            $s_count[ord($s[$i - $m]) - ord('a')]--;
            if ($p_count == $s_count) {
                $ans[] = $i - $m + 1;
            }
        }
        return $ans;
    }
}