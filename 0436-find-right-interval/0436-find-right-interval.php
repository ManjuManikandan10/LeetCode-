class Solution {
    function findRightInterval($intervals) {
        $n = count($intervals);
        for ($i = 0; $i < $n; $i++) {
            $intervals[$i][] = $i; // Add original index
        }
        sort($intervals);

        $ans = array_fill(0, $n, -1);
        for ($i = 0; $i < $n; $i++) {
            $target = $intervals[$i][1];
            $j = $this->binarySearch($intervals, $target);
            if ($j < $n) {
                $ans[$intervals[$i][2]] = $intervals[$j][2];
            }
        }

        return $ans;
    }

    private function binarySearch($intervals, $target) {
        $low = 0;
        $high = count($intervals) - 1;

        while ($low <= $high) {
            $mid = floor(($low + $high) / 2);
            if ($intervals[$mid][0] >= $target) {
                $high = $mid - 1;
            } else {
                $low = $mid + 1;
            }
        }

        return $low;
    }
}