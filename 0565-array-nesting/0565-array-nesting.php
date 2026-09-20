class Solution
{

    /**
     * @param Integer[] $nums
     *
     * @return Integer
     */
    function arrayNesting($nums)
    {
        $max = 1;

        $count = count($nums);
        for ($i = 0; $i < $count; $i++) {
            if ($nums[$i] >= 0) {
                $j = $i;
                $c = 0;
                while ($nums[$j] >= 0) {
                    $n = $nums[$j];
                    $nums[$j] = -1;
                    $j = $n;
                    $c++;
                }
                $max = max($max, $c);
            }
        }

        return $max;
    }
}