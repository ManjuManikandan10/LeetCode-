class Solution {

    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    function checkInclusion($s1, $s2) {

        $cnt = strlen($s2); 
        $len = strlen($s1);
        $hash = [];
        for($i=0; $i < $len; $i++) {
            if(!isset($hash[$s1[$i]])) $hash[$s1[$i]] = 0;
            $hash[$s1[$i]]++;
        }

        for($i=0; $i <= $cnt - $len; $i++) {
            $tmp = $hash;
            for($j=$i;$j<$i+$len;$j++){
                if(isset($tmp[$s2[$j]]) && $tmp[$s2[$j]]!==0) {
                    $tmp[$s2[$j]]--;
                }
            }
            
            if(array_sum($tmp) == 0) {
                return true;
            }
        }
        return false;
    }
}