func minTime(n int, edges [][]int, k int) int {
    g := make([][][]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], []int{e[1], e[2]})
        g[e[1]] = append(g[e[1]], []int{e[0], e[2]})
    }
    l, r := 0, math.MaxInt32
    for l < r { // binary search
        m := (l + r) / 2
        if cc(g, n, m) >= k { // if number of connected components > k, then we can decreae interval
            r = m
        } else { // else increase interval
            l = m + 1
        }
    }
    return l
}



func cc(g [][][]int, n, t int) int { // bfs returning number of connected components with (time > t); if lesser - edge is removed
    was := make([]bool, n)
    c := 0
    var bfs func(start int)
    bfs = func(start int) {
        was[start] = true
        q := [][]int{{start, 0}} // {vertex, time}
        for len(q) > 0 {
            nq := [][]int{}
            for k := range q {
                v := q[k][0] // current vertex
                for _, u := range g[v] {
                    if !was[u[0]] && u[1] > t { // if edge is not removed (it's time > allowed time), then it's connected component; else it is disconnected component
                        nq = append(nq, u)
                        was[u[0]] = true
                    }
                }
            }
            q = nq
        }
    }
    for v := 0; v < n; v++ { // iterate over all vertexes, counting number of connected components
        if !was[v] {
            c++
            bfs(v)
        }
    }
    return c
}