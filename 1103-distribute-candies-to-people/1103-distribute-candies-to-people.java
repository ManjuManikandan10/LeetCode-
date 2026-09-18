class Solution {
    public int[] distributeCandies(int candies, int num_people) {
        int[] ans = new int[num_people];
        int currCandyCount = 1;
        while(candies > 0){
            for(int i = 0; i < num_people; i++){
                if(candies >= currCandyCount){
                    ans[i] += currCandyCount;
                    candies -= currCandyCount;
                    currCandyCount++;
                }
                else{
                    ans[i] += candies;
                    candies = 0;
                }
            }
        }
        return ans;
    }
}