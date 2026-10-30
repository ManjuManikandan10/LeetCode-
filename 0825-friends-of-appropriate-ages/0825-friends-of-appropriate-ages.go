func numFriendRequests(ages []int) int {
	counts := make([]int, 121)
	for _, age := range ages {
		counts[age]++
	}
	var total int
	for ageX := 15; ageX <= 120; ageX++ {
		if counts[ageX] == 0 {
			continue
		}
		for ageY := int(0.5*float64(ageX)+7) + 1; ageY <= ageX; ageY++ {
			total += counts[ageX] * (counts[ageY] - btoi(ageX == ageY))
		}
	}
	return total
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}