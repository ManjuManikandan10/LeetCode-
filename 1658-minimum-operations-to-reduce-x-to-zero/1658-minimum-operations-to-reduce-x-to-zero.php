class Solution {
    public function minOperations($nums, $x) {
        $n = count($nums);
        $sum = array_sum($nums) - $x;

        if ($sum == 0) {
            return $n;
        }

        $start = 0;
        $curSum = 0;
        $len = 0;

        for ($end = 0; $end < $n; $end++) {
            $curSum += $nums[$end];

            while ($start < $n && $curSum > $sum) {
                $curSum -= $nums[$start];
                $start++;
            }

            if ($curSum == $sum) {
                $len = max($len, $end - $start + 1);
            }
        }

        if ($len == 0) {
            return -1;
        }

        return $n - $len;
    }
}