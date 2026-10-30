func regionsBySlashes(grid []string) int {
    g := gridToMatrix(grid)
    r := 0
    for j := 0; j < len(g); j++ {
        for i := 0; i < len(g[0]); i++ {
            if g[j][i] == 0 { // got 0 -> fill with -1 and increase result
                r ++
                scan(g, j, i)
            }
        }
    }
    return r
}

func scan(g [][]int, y, x int) { // bfs for grid moving, if we have connected 0 area, fill it with -1's
    d := []int{-1, 0, 1, 0, -1}
    q := [][]int{{y, x}}
    g[y][x] = -1
    for len(q) > 0 {
        nq := [][]int{}
        for k := range q {
            for i := 0; i <= 3; i++ {
                dy, dx := q[k][0] + d[i], q[k][1] + d[i+1]
                if dy >= 0 && dy < len(g) && dx >= 0 && dx < len(g[0]) && g[dy][dx] == 0 {
                    g[dy][dx] = -1
                    nq = append(nq, []int{dy, dx})
                }
            }
        }
        q = nq
    }
}


// helper func example
// '/'    '\\'     
// 001    100
// 010    010
// 100    001
func gridToMatrix(grid []string) [][]int { // helper func to make '/\\ ' stuff to 0 and 1 matrix. every grid cell -> 3x3 square
	n := len(grid)
	s := n * 3
	g := make([][]int, s)
	for i := range g {
		g[i] = make([]int, s)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r := i * 3
			c := j * 3
			switch grid[i][j] {
			case '/':
				g[r][c+2] = 1
				g[r+1][c+1] = 1
				g[r+2][c] = 1
			case '\\':
				g[r][c] = 1
				g[r+1][c+1] = 1
				g[r+2][c+2] = 1
			}
		}
	}
	return g
}