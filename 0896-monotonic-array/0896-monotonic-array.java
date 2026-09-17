class Solution {
    public boolean isMonotonic(int[] nums) {
        boolean isIncreasing = true;
        boolean isDecreasing = true;
        for (int i =1; i<nums.length; i++){
            if (nums[i]>nums[i-1]){
                isDecreasing = false;
            }
            else if (nums[i]<nums[i-1]){
                isIncreasing = false;
            }
            if(!isIncreasing && !isDecreasing){
                return false;
            }
        }
        return true;
    }
}