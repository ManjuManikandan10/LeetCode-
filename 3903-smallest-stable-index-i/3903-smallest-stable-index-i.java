class Solution {
    public int firstStableIndex(int[] nums, int k) {
        int n = nums.length;

        int[] maxarr = new int[n];
        maxarr[0] = nums[0];
        int max = nums[0];

        for (int i = 1; i < n; i++) {
            if (max < nums[i]) {
                maxarr[i] = nums[i];
                max = nums[i];
            } else {
                maxarr[i] = max;
            }
        }

        int[] minarr = new int[n];
        minarr[n - 1] = nums[n - 1];
        int min = nums[n - 1];

        for (int i = n - 2; i >= 0; i--) {
            if (min > nums[i]) {
                minarr[i] = nums[i];
                min = nums[i];
            } else {
                minarr[i] = min;
            }
        }

        for (int i = 0; i < n; i++) {
            int res = maxarr[i] - minarr[i];
            if (res <= k) {
                return i;
            }
        }

        return -1;
    }
}