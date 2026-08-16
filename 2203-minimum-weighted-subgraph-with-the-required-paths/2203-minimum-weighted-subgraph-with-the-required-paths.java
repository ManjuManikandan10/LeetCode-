class Solution {
    public long minimumWeight(int n, int[][] edges, int src1, int src2, int dest) {
        List<List<int[]>> adjList = new ArrayList<>(n), revAdjList = new ArrayList<>(n);
        for(int i=0; i<n; ++i) {
            adjList.add(new ArrayList<int[]>());
            revAdjList.add(new ArrayList<int[]>());
        }
        for(int[] edge:edges) {
            adjList.get(edge[0]).add(new int[]{edge[1], edge[2]});
            revAdjList.get(edge[1]).add(new int[]{edge[0], edge[2]});
        }

        long[] arr1 = singleSourceShortestPathDijkstra(adjList, n, src1);
        long[] arr2 = singleSourceShortestPathDijkstra(adjList, n, src2);
        long[] arr3 = singleSourceShortestPathDijkstra(revAdjList, n, dest);
        long minCost = Long.MAX_VALUE;

        for(int i=0; i<n; ++i) {
            long src1_x = arr1[i], src2_x = arr2[i], x_dest = arr3[i];
            if(src1_x == -1 || src2_x == -1 || x_dest == -1) {
                continue;
            }
            minCost = Math.min(minCost, src1_x + src2_x + x_dest);
        }
        if(minCost == Long.MAX_VALUE) return -1;

        return minCost;
    }

    private long[] singleSourceShortestPathDijkstra(List<List<int[]>> adjList, int n, int src) {
        PriorityQueue<long[]> minHeap = new PriorityQueue<long[]>((a, b) -> Long.compare(a[1], b[1]));
        minHeap.add(new long[]{src, 0});
        long[] dist = new long[n];
        Arrays.fill(dist, Long.MAX_VALUE);
        dist[src] = 0;
        while(!minHeap.isEmpty()) {
            long[] top = minHeap.poll();
            int node = (int)top[0];
            if(top[1] > dist[node]) continue;
            for(int[] next:adjList.get(node)) {
                if(dist[node] < Long.MAX_VALUE && dist[node] + next[1] < dist[next[0]]) {
                    dist[next[0]] = dist[node] + next[1];
                    minHeap.add(new long[]{next[0], dist[node] + next[1]});
                }
            }
        }
        for(int i=0; i<n; ++i) {
            if(dist[i] == Long.MAX_VALUE) {
                dist[i] = -1;
            }
        }
        return dist;
    }
}