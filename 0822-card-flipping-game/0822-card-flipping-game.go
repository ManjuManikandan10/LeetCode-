func flipgame(fronts []int, backs []int) int {
	impossible := make([]bool, 2001)
	exists := make([]bool, 2001)
	for i := range fronts {
		front := fronts[i]
		back := backs[i]

		exists[front] = true
		exists[back] = true

		if front == back {
			impossible[front] = true
		}
	}

	for i := 1; i < len(impossible); i++ {
		if exists[i] && !impossible[i] {
			return i
		}
	}

	return 0
}