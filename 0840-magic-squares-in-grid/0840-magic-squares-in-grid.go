func numMagicSquaresInside(grid [][]int) int {
    r, c := len(grid), len(grid[0])
    if r < 3 || c < 3 {
        return 0
    }

    ans := 0
    for i := 0; i+2 < r; i++ {
        for j := 0; j+2 < c; j++ {
            used := make([]bool, 10)
            ok := true

            for x := 0; x < 3 && ok; x++ {
                for y := 0; y < 3; y++ {
                    v := grid[i+x][j+y]
                    if v < 1 || v > 9 || used[v] {
                        ok = false
                        break
                    }
                    used[v] = true
                }
            }
            if !ok {
                continue
            }

            s := grid[i][j] + grid[i][j+1] + grid[i][j+2]
            for x := 0; x < 3; x++ {
                if grid[i+x][j]+grid[i+x][j+1]+grid[i+x][j+2] != s {
                    ok = false
                }
            }
            for y := 0; y < 3; y++ {
                if grid[i][j+y]+grid[i+1][j+y]+grid[i+2][j+y] != s {
                    ok = false
                }
            }
            if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != s {
                ok = false
            }
            if grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] != s {
                ok = false
            }

            if ok {
                ans++
            }
        }
    }
    return ans
}