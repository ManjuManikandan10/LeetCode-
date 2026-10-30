
func bestRotation(nums []int) int {
	dp := make([]int, len(nums))
	rs := 0
	mx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= i {
			mx++
		}
		if nums[i] > 0 && nums[i] < len(nums) {
			if nums[i] > i {
				dp[i+1] += 1
				k := i + len(nums) - nums[i] + 1
				if k < len(nums) {
					dp[k] -= 1
				}
			} else {
				dp[i-nums[i]+1] -= 1
				if i+1 < len(nums) {
					dp[i+1] += 1
				}
			}
		}
	}
	dp[0] = mx
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i]
		if dp[i] > mx {
			mx = dp[i]
			rs = i
		}
	}
	return rs
}