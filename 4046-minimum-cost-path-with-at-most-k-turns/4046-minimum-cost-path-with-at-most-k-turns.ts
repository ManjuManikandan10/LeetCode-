function minCost(grid: number[][], k: number): number {
    const n = grid.length, m = grid[0].length;
    const dirs = [0, 1, 0, -1, 0];

    // dp[y][x][turns][direction]
    const dp = new Array(n)
        .fill(0).map(() => new Array(m)
            .fill(0).map(() => new Array(k + 1)
                .fill(0).map(() => new Array(4).fill(Number.MAX_SAFE_INTEGER))
            )
        );

    // [cost, turns, x, y, current direction]
    // min-heap
    const pq = new PriorityQueue<number[]>(([a], [b]) => a - b);

    // -1 because initial turn doesn't count
    pq.push([grid[0][0], -1, 0, 0, -1]);

    // 1 * 1 matrix edge case
    dp[0][0][0][0] = grid[0][0];

    while (!pq.isEmpty()) {
        const [curr_cost, curr_turns, x, y, curr_dir] = pq.pop();

        // Path found
        if (x == m - 1 && y == n - 1) return curr_cost;

        // More optimal way found
        if (curr_turns != -1 && dp[y][x][curr_turns][curr_dir] < curr_cost) continue;

        for (let i = 0; i < 4; i++) {
            const nx = x + dirs[i];
            const ny = y + dirs[i + 1];

            if (nx < 0 || ny < 0 || nx >= m || ny >= n) continue;

            let next_turns = curr_turns;
            if (curr_dir != i) next_turns++;
            // Too many turns
            if (next_turns > k) continue;

            const next_cost = curr_cost + grid[ny][nx];
            // More optimal path visited
            if (dp[ny][nx][next_turns][i] <= next_cost) continue;
            dp[ny][nx][next_turns][i] = next_cost;

            pq.push([next_cost, next_turns, nx, ny, i]);
        }
    }

    return -1;
};