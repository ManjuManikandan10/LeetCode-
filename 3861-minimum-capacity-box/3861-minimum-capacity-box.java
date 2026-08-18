class Solution {
    public int minimumIndex(int[] arr, int k) {
        int ans = -1;
        int min = -1;

        for (int i = 0; i < arr.length; i++) {
            if (arr[i] >= k) {
                if (min == -1 || arr[i] < min) {
                    min = arr[i];
                    ans = i;
                }
            }
        }

        return ans;
    }
}