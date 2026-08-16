class Solution {

    int max = 0;
    public int maximalPathQuality(int[] Values, int[][] edges, int maXTime) {
        int n = Values.length;

        List<int[]>[] graph = new ArrayList[n];
        for (int i = 0; i < n; i++) graph[i] = new ArrayList<>();

        for (int[] e : edges) {
            graph[e[0]].add(new int[]{e[1], e[2]});
            graph[e[1]].add(new int[]{e[0], e[2]});
        }

        int[] freq = new int[n];
        dfs(0, maXTime, 0, Values, graph, freq);

        return max;
    }

    private void dfs(int node, int time, int score, int[] values,
                     List<int[]>[] graph, int[] freq) {

        if (freq[node] == 0) {
            score += values[node];
        }

        freq[node]++;

        if (node == 0) {
            max = Math.max(max, score);
        }

        for (int[] nei : graph[node]) {
            int next = nei[0];
            int cost = nei[1];

            if (time >= cost) {
                dfs(next, time - cost, score, values, graph, freq);
            }
        }

        freq[node]--;
    }
}