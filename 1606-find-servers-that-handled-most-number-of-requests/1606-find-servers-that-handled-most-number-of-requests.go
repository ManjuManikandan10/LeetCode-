func busiestServers(k int, arrival []int, load []int) []int {
	var jobs, serversBef, serversAfter PriorityQueue
	var reqs = make([]int, k)
	for i := 0; i < k; i++ {
		var node = &Item{
			num:      i,
			priority: i,
			index:    i,
		}

		serversBef = append(serversBef, node)
	}
	heap.Init(&serversBef)

	for i := range arrival {
		sec := arrival[i]
		load := load[i]
		num := i % k
		if num == 0 {
			serversAfter = serversBef
			serversBef = PriorityQueue{}
		}

		for jobs.Len() > 0 {
			job := heap.Pop(&jobs).(*Item)
			if job.priority > sec {
				heap.Push(&jobs, job)
				break
			}
			node := &Item{num: job.num, priority: job.num}
			if job.num < num {
				heap.Push(&serversBef, node)
			} else {
				heap.Push(&serversAfter, node)
			}
		}

		var node *Item
		if serversAfter.Len() > 0 {
			node = heap.Pop(&serversAfter).(*Item)
		} else if serversBef.Len() > 0 {
			node = heap.Pop(&serversBef).(*Item)
		} else {
			continue
		}
		job := &Item{
			num:      node.num,
			priority: sec + load,
		}
		heap.Push(&jobs, job)
		reqs[node.num]++
	}
	var mr = map[int][]int{}
	var max = 0
	for i, r := range reqs {
		if r > max {
			max = r
		}
		mr[r] = append(mr[r], i)
	}

	return mr[max]
}

type Item struct {
	num      int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].num < pq[j].num
	}
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}