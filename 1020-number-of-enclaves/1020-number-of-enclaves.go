func numEnclaves(grid [][]int) int {
    ans := 0
    
    // Tranvese land from the boundary to inside of the grid, 
    // Calculate the connected adjancent land cell
    for i := 1; i < len(grid) - 1; i++ {
         if grid[i][0] == 1 {
              dfs(grid, i, 0)
         }

         if grid[i][len(grid[0])-1] == 1 {
            dfs(grid, i, len(grid[0]) - 1)
         }     
    }

    for j := 1; j < len(grid[0]) - 1; j++ {
        if grid[0][j] == 1 {
            dfs(grid, 0, j)
        }

        if grid[len(grid)-1][j] == 1 {
            dfs(grid, len(grid)-1, j)
        }
          
    }

    // Just counting the answer
    for i := 1; i < len(grid) - 1; i++ {
        for j := 1; j < len(grid[0]) - 1; j++ {
            if grid[i][j] == 1 {
                ans++
            }
        }
    }


    return ans
}


func dfs(grid [][]int, i, j int ) {
    // out of boundary 
    if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[i]) -1 {
        return
    }

    // Is water, of course 
    if grid[i][j] != 1 {
        return
    }

    // make land was tranversed
    grid[i][j] = 2

    // tranverse adjacent land
    dfs(grid, i, j-1)
    dfs(grid, i,j+1)
    dfs(grid, i+1,j)
    dfs(grid,i-1,j)
}