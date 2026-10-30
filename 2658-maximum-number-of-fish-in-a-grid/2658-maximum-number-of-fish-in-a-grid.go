func findMaxFish(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }


    var dfs func(x, y int) int
    dfs = func(x, y int) int {
        if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 || visited[x][y] {
            return 0
        }

        visited[x][y] = true
        currSum := grid[x][y] 

        for _, d := range dir {
            nx, ny := x+d[0], y+d[1]
            currSum += dfs(nx, ny)
        }
        return currSum
    }

    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 && !visited[i][j] {
                currSum := dfs(i, j)
                if currSum > ans {
                    ans = currSum
                }
            }
        }
    }
    return ans
}