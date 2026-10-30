func find(x int, parent []int) int {
	for x != parent[x] {
		parent[x] = parent[parent[x]]
		x = parent[x]
	}
	return x
}

func union(a, b int, parent, size []int) {
	ra := find(a, parent)
	rb := find(b, parent)
	if ra == rb {
		return
	}
	if size[ra] < size[rb] {
		ra, rb = rb, ra
	}
	parent[rb] = ra
	size[ra] += size[rb]
}

func maxAlternatingSum(nums []int, swaps [][]int) int64 {
	n := len(nums)
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	for _, e := range swaps {
		union(e[0], e[1], parent, size)
	}

	components := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		r := find(i, parent)
		components[r] = append(components[r], i)
	}

	var res int64
	for _, idxs := range components {
		vals := make([]int64, 0, len(idxs))
		cntEven := 0
		for _, i := range idxs {
			vals = append(vals, int64(nums[i]))
			if i%2 == 0 {
				cntEven++
			}
		}
		sort.Slice(vals, func(i, j int) bool { return vals[i] > vals[j] })
		var pos, neg int64
		for i := 0; i < cntEven && i < len(vals); i++ {
			pos += vals[i]
		}
		for i := cntEven; i < len(vals); i++ {
			neg += vals[i]
		}
		res += pos - neg
	}
	return res
}