class NumArray {
    private int[] pre;
    public NumArray(int[] nums) {
        int n = nums.length;
        pre = new int[n+1];

        pre[0] = 0;
        for(int i = 1;i<=n;i++){
            pre[i] = pre[i-1] + nums[i-1];
        }    
    }
    
    public int sumRange(int left, int right) {
        return pre[right+1] - pre[left];
    }
}