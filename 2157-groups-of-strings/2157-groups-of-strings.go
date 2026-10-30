func groupStrings(words []string) []int {
	getBits := func(w string) int {
		var res int
		for _, ch := range w {
			res |= 1 << int(ch-'a')
		}
		return res
	}

	n := len(words)

	// DSU setup
	parent := make([]int, n)
	size := make([]int, n)
	for i := range words {
		parent[i] = i
		size[i] = 1
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		parent[a] = find(parent[a]) // path compression
		return parent[a]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = ra
			size[ra] += size[rb]
		}
	}

	// For each word, parse its bit-masks and add to a list of indices per mask
	masks := make(map[int][]int, n)
	for i, w := range words {
		bits := getBits(w)
		masks[bits+1<<26] = append(masks[bits+1<<26], i) // add wildcard
		for b := 1; b < 1<<27; b <<= 1 {
			if bits&b > 0 {
				replaced := (bits &^ b) | 1<<26
				masks[replaced] = append(masks[replaced], i) // replacement wildcard
			}
		}
	}

	// For each mask, union matched indices
	for _, indices := range masks {
		for i := 0; i < len(indices)-1; i++ {
			union(indices[i], indices[i+1])
		}
	}

	// Collect group sizes
	groupSize := make(map[int]int)
	var largestGroup int
	for i := range words {
		r := find(i)
		groupSize[r] = size[r]
		if groupSize[r] > largestGroup {
			largestGroup = groupSize[r]
		}
	}

	res := []int{len(groupSize), largestGroup}
	return res
}