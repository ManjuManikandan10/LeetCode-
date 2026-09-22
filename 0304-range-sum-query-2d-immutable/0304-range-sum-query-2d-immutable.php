class NumMatrix {

    private array $matrix = [];
    /**
     * @param Integer[][] $matrix
     */
    function __construct($matrix) {
        for($i = 0;$i < count($matrix[0]);$i++)
        {
            $matrix[0][$i] = $matrix[0][$i] + $matrix[0][$i - 1];
        }

        for($j = 1;$j < count($matrix);$j++)
        {
            $matrix[$j][0] = $matrix[$j][0] + $matrix[$j - 1][0];
        }


        for($j = 1;$j < count($matrix);$j++)
        {
            for($t = 1;$t < count($matrix[0]);$t++)
            {
                $matrix[$j][$t] = $matrix[$j][$t] + $matrix[$j-1][$t] + $matrix[$j][$t-1] - $matrix[$j-1][$t - 1];
            }  
        }

        $this->matrix = $matrix;
    }
  
    /**
     * @param Integer $row1
     * @param Integer $col1
     * @param Integer $row2
     * @param Integer $col2
     * @return Integer
     */
    function sumRegion($row1, $col1, $row2, $col2) {
       return $this->matrix[$row2][$col2] - $this->matrix[$row1 - 1][$col2] - $this->matrix[$row2][$col1 - 1] + $this->matrix[$row1 - 1][$col1 - 1];
    }
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * $obj = NumMatrix($matrix);
 * $ret_1 = $obj->sumRegion($row1, $col1, $row2, $col2);
 */