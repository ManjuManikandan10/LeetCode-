func reachNumber(target int) int {
	var sum, step int
	if target < 0 {
		target = -target
	}

	for {
		sum += step
		if sum >= target {
			break
		}
		step++
	}

	if (sum-target)&1 == 0 {
		return step
	}

	if step&1 == 0 {
		return step + 1
	}
	return step + 2
}