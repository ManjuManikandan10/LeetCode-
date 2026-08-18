class Solution {
    public int minAbsoluteDifference(int[] nums) {
        int min = Integer.MAX_VALUE;
        for (int a = 0; a < nums.length - 1; a++) {
            for (int b = a + 1; b < nums.length; b++) {
                if ((nums[a] == 1 && nums[b] == 2) || (nums[a] == 2 && nums[b] == 1)) {
                    min = Math.min(min, Math.abs(a-b));
                }
            }
        }
        return min == Integer.MAX_VALUE ? -1 : min;
    }
}