// Complexity:
// Time O(NLogN) and Space O(N) where N is the length of
// intervals.
func intersectionSizeTwo(intervals [][]int) int {
	// The most recent two values of our greedy nums
	recentTwo := []int{-2, -1}
	result := 0

	sortedIntervals := slices.Clone(intervals)
	slices.SortFunc(sortedIntervals, func(i1, i2 []int) int {
		return i1[1] - i2[1]
	})

	for _, interval := range sortedIntervals {
		if recentTwo[1] < interval[0] {
			recentTwo[0] = interval[1] - 1
			recentTwo[1] = interval[1]
			result += 2
		} else if recentTwo[0] < interval[0] {
			recentTwo[0] = min(interval[1]-1, recentTwo[1])
			recentTwo[1] = interval[1]
			result += 1
		}
	}
	return result
}