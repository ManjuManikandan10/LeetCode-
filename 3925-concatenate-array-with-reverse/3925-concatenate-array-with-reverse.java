class Solution {
    public int[] concatWithReverse(int[] nums) {
        int [] sol = Arrays.copyOf(nums,nums.length*2);
        int index = nums.length;
        for(int i = index-1; i >= 0; i--){
            sol[index++] = nums[i];
        }
         return sol;
    }
}