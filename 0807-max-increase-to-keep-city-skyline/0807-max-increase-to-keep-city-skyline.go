func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)

	for r, row := range grid {
		for c, h := range row {
			rowMax[r] = max(rowMax[r], h)
			colMax[c] = max(colMax[c], h)
		}
	}

	total := 0
	for r, row := range grid {
		for c, h := range row {
			total += min(rowMax[r], colMax[c]) - h
		}
	}

	return total
}