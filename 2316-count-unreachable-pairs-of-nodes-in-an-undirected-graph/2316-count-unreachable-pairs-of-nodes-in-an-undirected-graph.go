func countPairs(n int, edges [][]int) int64 {
    parents := make([]int, n)
    ranks := make([]int, n)

    for i := 0; i < n; i++ {
        parents[i] = i
    }

    for _, edge := range edges {
        union(edge[0], edge[1], parents, ranks)
    }

    groups := make(map[int]int)
    for i := 0; i < n; i++ {
        parent := find(i, parents)
        groups[parent]++
    }

    unReachableConns := 0
    totalRemainingConns := n
    for _, groupSize := range groups {
        unReachableConns += groupSize * (totalRemainingConns - groupSize)
        totalRemainingConns -= groupSize
    }

    return int64(unReachableConns)
}

func union(idx1, idx2 int, parents, ranks []int) {
    p1 := find(idx1, parents)
    p2 := find(idx2, parents)

    if p1 == p2 {
        return
    }

    if ranks[p1] > ranks[p2] {
        parents[p2] = p1
    } else if ranks[p1] < ranks[p2] {
        parents[p1] = p2
    } else {
        parents[p2] = p1
        ranks[p1]++
    }
}

func find(idx int, parents []int) int {
    if idx != parents[idx] {
        parents[idx] = find(parents[idx], parents)
    }
    return parents[idx]
}