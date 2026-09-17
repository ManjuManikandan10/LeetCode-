class Solution {
    public int smallestRangeI(int[] nums, int k) {
        int maxi = Integer.MIN_VALUE;
        int mini = Integer.MAX_VALUE;
        int n = nums.length;
        for(int i=0;i<n;i++)
            {
                maxi = Math.max(nums[i],maxi);
                mini = Math.min(nums[i],mini);
            }
        int ans = Math.abs(maxi-mini)-2*k;
        if(ans<0)ans=0;
        return ans;
    }
}