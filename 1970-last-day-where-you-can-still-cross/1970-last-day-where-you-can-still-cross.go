func latestDayToCross(m int, n int, cells [][]int) int {
    res := 0
    lo, hi := 0, len(cells)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        matrix := getMatrix(cells, mid, m, n)
        if bfsReachable(matrix) {
            res = mid + 1
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return res
}

func bfsReachable(matrix [][]int) bool {
    frontier := make([][2]int, 0)
    visited := map[[2]int]bool{}
    for j := range matrix[0] {
        if matrix[0][j] == 0 {
            frontier = append(frontier, [2]int{0, j})
            visited[[2]int{0, j}] = true
        }
    }

    for len(frontier) > 0 {
        var next [][2]int
        for _, node := range frontier {
            for _, dir := range [][]int{{0,1},{0,-1},{1,0},{-1,0}} {
                i, j := node[0]+dir[0], node[1]+dir[1]
                pair := [2]int{i, j}
                if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) ||
                   visited[pair] || matrix[i][j] == 1 {
                    continue
                }
                if i == len(matrix)-1 {
                    return true
                }
                visited[pair] = true
                next = append(next, pair)
            }
        }
        frontier = next
    }
    return false
}

func getMatrix(cells [][]int, mid int, m int, n int) [][]int {
    matrix := make([][]int, m)
    for i := 0; i < m; i++ {
        matrix[i] = make([]int, n)
    }
    for i := 0; i <= mid; i++ {
        r, c := cells[i][0]-1, cells[i][1]-1
        matrix[r][c] = 1
    }
    return matrix
}