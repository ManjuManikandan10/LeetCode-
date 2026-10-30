type TreapNode struct {
	Start int
	End   int
	Sum   int64
	Left  *TreapNode
	Right *TreapNode
}

func merge(a, b *TreapNode) *TreapNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if a.Sum > b.Sum {
		a.Right = merge(a.Right, b)
		return a
	}

	b.Left = merge(a, b.Left)
	return b
}

func split(a *TreapNode, start int) (*TreapNode, *TreapNode) {
	if a == nil {
		return nil, nil
	}

	if a.Start < start {
		u, v := split(a.Right, start)
		a.Right = u
		return a, v
	}

	u, v := split(a.Left, start)
	a.Left = v
	return u, a
}

func insert(a, b *TreapNode) *TreapNode {
	if a == nil || a.Sum < b.Sum {
		b.Left, b.Right = split(a, b.Start)
		return b
	}

	if a.Start < b.Start {
		a.Right = insert(a.Right, b)
	} else {
		a.Left = insert(a.Left, b)
	}

	return a
}

func remove(a *TreapNode, start int) (*TreapNode, *TreapNode) {
	var p, b *TreapNode

	b = a
	for {
		if b.Start <= start && start < b.End {
			break
		}

		if start < b.Start {
			p, b = b, b.Left
		} else {
			p, b = b, b.Right
		}
	}

	if p == nil {
		a = merge(b.Left, b.Right)
	} else if p.Left == b {
		p.Left = merge(b.Left, b.Right)
	} else {
		p.Right = merge(b.Left, b.Right)
	}

	return a, b
}

func maximumSegmentSum(nums, queries []int) []int64 {
	length := len(nums)

	prefixSum := make([]int64, len(nums)+1)
	for i, num := range nums {
		prefixSum[i+1] = prefixSum[i] + int64(num)
	}

	treap := &TreapNode{
		Start: 0,
		End:   length,
		Sum:   prefixSum[length],
	}

	res := make([]int64, 0, length)
	for _, query := range queries {
        var a *TreapNode
        treap, a = remove(treap, query)

		if a.Start != query {
			treap = insert(treap, &TreapNode{
				Start: a.Start,
				End:   query,
				Sum:   prefixSum[query] - prefixSum[a.Start],
			})
		}

		if query+1 != a.End {
			treap = insert(treap, &TreapNode{
				Start: query + 1,
				End:   a.End,
				Sum:   prefixSum[a.End] - prefixSum[query+1],
			})
		}

		if treap == nil {
			res = append(res, 0)
		} else {
			res = append(res, treap.Sum)
		}
	}

	return res
}
