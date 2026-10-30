
func maxPoints(grid [][]int, queries []int) []int {
	rs := make([]int, len(queries))
	queryQueue := make([]int, len(queries))
	for i := range queries {
		queryQueue[i] = i
	}
	sort.Slice(queryQueue, func(i, j int) bool {
		return queries[queryQueue[i]] < queries[queryQueue[j]]
	})
	r, c := len(grid), len(grid[0])
	checkProcess := make([]int, r*c)
	lastPosition := &IntHeap{
		data: []int{},
		grid: grid,
	}
	heap.Init(lastPosition)
	heap.Push(lastPosition, 0)
	lastVal := 0
	for len(queryQueue) > 0 {
		idxQueries := queryQueue[0]
		val := queries[idxQueries]
		queryQueue = queryQueue[1:]
		newVal := 0
		for lastPosition.Len() > 0 {
			vl := heap.Pop(lastPosition).(int)
			ri, ci := vl/c, vl%c
			if grid[ri][ci] >= val {
				heap.Push(lastPosition, vl)
				break
			} else {
				if checkProcess[vl] == 0 {
					newVal++
				} else {
					continue
				}
				checkProcess[vl] = 1
				//check and push 4 new point
				if ri-1 >= 0 && checkProcess[(ri-1)*c+ci] == 0 {
					heap.Push(lastPosition, (ri-1)*c+ci)
				}
				if ri+1 < r && checkProcess[(ri+1)*c+ci] == 0 {
					heap.Push(lastPosition, (ri+1)*c+ci)
				}
				//check and push 4 new point
				if ci-1 >= 0 && checkProcess[ri*c+ci-1] == 0 {
					heap.Push(lastPosition, ri*c+ci-1)
				}
				if ci+1 < c && checkProcess[ri*c+ci+1] == 0 {
					heap.Push(lastPosition, ri*c+ci+1)
				}
			}
		}
		lastVal += newVal
		rs[idxQueries] = lastVal
	}
	return rs
}

type IntHeap struct {
	data []int
	grid [][]int
}

func (h IntHeap) Len() int { return len(h.data) }
func (h IntHeap) Less(i, j int) bool {
	ci := h.data[i] % len(h.grid[0])
	ri := h.data[i] / len(h.grid[0])
	cj := h.data[j] % len(h.grid[0])
	rj := h.data[j] / len(h.grid[0])
	return h.grid[ri][ci] < h.grid[rj][cj]
}
func (h IntHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}