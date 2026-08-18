class Solution {
    public int minimumSwaps(int[] nums) {
        int count = 0;
        for(int i = 0;i < nums.length;i++){
            if(nums[i] == 0){
                count++;
            }
        }
        if(count == 0){
            return 0;
        }
        int x = 0;
        for(int i = nums.length - 1;i >= (nums.length - count);i--){
            if(nums[i] == 0){
                x++;
            }
        }
        return (count - x);
    }
}