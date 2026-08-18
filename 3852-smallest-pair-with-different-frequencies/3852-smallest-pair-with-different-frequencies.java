class Solution {
    public int[] minDistinctFreqPair(int[] nums) {
        int freq[]=new int[101];
        Arrays.sort(nums);
        for(int i:nums){
            freq[i]++;
        }
        int min=nums[0];
        int minfreq=freq[min];
        for(int i:nums){
            if(i!=min&&minfreq!=freq[i]){
                return new int[]{min,i};
            }
        }
        return new int[]{-1,-1};
        
    }
}