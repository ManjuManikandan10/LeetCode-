type UnionFind struct {
	Parents []int
	Ranks   []int
}

func newUF(size int) *UnionFind {
	parents := make([]int, 0)
	for n := range size {
		parents = append(parents, n)
	}

	ranks := make([]int, size)

	return &UnionFind{
		Parents: parents,
		Ranks:   ranks,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parents[x] != x {
		uf.Parents[x] = uf.Find(uf.Parents[x])
	}
	return uf.Parents[x]
}


func (uf *UnionFind) Union(x, y int) {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)

	if uf.Ranks[xRoot] > uf.Ranks[yRoot] {
		uf.Parents[yRoot] = xRoot
	} else {
		uf.Parents[xRoot] = yRoot
		if uf.Ranks[xRoot] == uf.Ranks[yRoot] {
			uf.Ranks[yRoot]++
		}
	}
}

func makeConnected(n int, connections [][]int) int {
    uf := newUF(n)
    potentialWires := 0
    unconnectedComputers := n - 1
    for _, edge := range connections {
        u := edge[0]
        v := edge[1]

        if uf.Find(u) == uf.Find(v) {
            potentialWires++
        } else {
            uf.Union(u, v)
            unconnectedComputers--
        }
    }

    if potentialWires >= unconnectedComputers {
        return unconnectedComputers
    } else {
        return -1
    }
}