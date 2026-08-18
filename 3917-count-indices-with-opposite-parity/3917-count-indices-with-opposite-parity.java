class Solution {
    public int[] countOppositeParity(int[] nums) {
        int n=nums.length;
        int oddcount=0;
        int evecount=0;
        int res[]=new int[n];
        for(int i=n-1;i>=0;i--){
            if(nums[i]%2==0){
                evecount++;
                res[i]=oddcount;
            }
            else{
                oddcount++;
                res[i]=evecount;
            }
        }
        return res;
    }
}