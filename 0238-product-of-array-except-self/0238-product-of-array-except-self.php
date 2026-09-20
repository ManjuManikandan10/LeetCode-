class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    public function productExceptSelf($nums) {
        $n = count($nums);
        $output = array_fill(0, $n, 1);

        // Calculate left products
        $leftProduct = 1;
        for ($i = 0; $i < $n; $i++) {
            $output[$i] *= $leftProduct;
            $leftProduct *= $nums[$i];
        }

        // Calculate right products
        $rightProduct = 1;
        for ($i = $n - 1; $i >= 0; $i--) {
            $output[$i] *= $rightProduct;
            $rightProduct *= $nums[$i];
        }

        return $output;
    }
}