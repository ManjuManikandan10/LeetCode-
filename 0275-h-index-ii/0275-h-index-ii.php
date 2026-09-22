class Solution {

    /**
     * @param Integer[] $citations
     * @return Integer
     */
    function hIndex($citations) {
        $n = count($citations);
        $left = 0;
        $right = $n - 1;
        
        while ($left <= $right) {
            $mid = $left + intval(($right - $left) / 2);
            $papers = $n - $mid;
            
            if ($citations[$mid] == $papers) {
                return $papers;
            } elseif ($citations[$mid] < $papers) {
                $left = $mid + 1;
            } else {
                $right = $mid - 1;
            }
        }
        
        return $n - $left;
    }
}