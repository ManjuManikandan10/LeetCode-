class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @param Integer[] $nums3
     * @param Integer[] $nums4
     * @return Integer
     */
    function fourSumCount($nums1, $nums2, $nums3, $nums4) {
        $count = 0;
        $hash = [];
        
        foreach($nums1 as $v1){
            foreach($nums2 as $v2){
                $hash[$v1+$v2] = isset($hash[$v1+$v2]) ? $hash[$v1+$v2]+1 : 1;
            }
        }

        foreach($nums3 as $v3){
            foreach($nums4 as $v4){
                $count += $hash[-($v3+$v4)] ?? 0;
            }
        }


        return $count;
    }
}