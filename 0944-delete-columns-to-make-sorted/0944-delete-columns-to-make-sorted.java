class Solution {
    public int minDeletionSize(String[] strs) {

        int ans=0;
        int n=strs[0].length();
        int k=strs.length;
        for(int i=0;i<n;i++){
            // Iterate through each element in the column
            for(int j=1;j<k;j++)
            {
    // If the current element is lexicographically smaller than the previous element,
                // increment the delete count and break out of the loop
             if (strs[j].charAt(i) < strs[j - 1].charAt(i))
               {ans++;
               break;}
            }
        }
        return ans;
    }
}