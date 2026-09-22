class Solution {

    /**
     * @param Integer[][] $matrix
     * @param Integer $k
     * @return Integer
     */
    function kthSmallest($matrix, $k) {
        $vector = [];
        foreach($matrix as $row) {
            foreach ($row as $value) {
               $vector[] = $value;
            }
        }
        sort($vector);

        return $vector[$k - 1];
    }
}