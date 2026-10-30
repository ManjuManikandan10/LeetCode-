type Union struct {
    Parent []int
    Rank []int
}

func (u *Union) find(x int) int {
    if u.Parent[x] != x {
        u.Parent[x] = u.find(u.Parent[x])
    }
    return u.Parent[x]
}

func (u *Union) union(x, y int) bool {
    rootX := u.find(x)
    rootY := u.find(y)
    if rootX == rootY {
        return false
    }

    if u.Rank[rootX] < u.Rank[rootY] {
        u.Parent[rootX] = rootY
    } else if u.Rank[rootX] > u.Rank[rootY] {
        u.Parent[rootY] = rootX
    } else {
        u.Parent[rootY] = rootX
        u.Rank[rootX] += 1
    }

    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func minCost(n int, edges [][]int, k int) int {
    if len(edges) == 0 {
        return 0
    }
    
    sort.Slice(edges, func(i, j int) bool {
        return edges[i][2] < edges[j][2]
    })

    uf := &Union{
        Parent: make([]int, n),
        Rank: make([]int, n),
    }
    for i :=  range uf.Parent {
        uf.Parent[i] = i
    }

    mst := []int{}
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        if uf.union(u, v) {
           mst = append(mst, w) 
        }
    }

    sort.Sort(sort.Reverse(sort.IntSlice(mst)))
    if len(mst) < k-1 {
        return 0
    }

    if k == 1 {
        return mst[0]
    }

    mst = mst[k-1:]
    if len(mst) == 0 {
        return 0
    }

    return mst[0]
}