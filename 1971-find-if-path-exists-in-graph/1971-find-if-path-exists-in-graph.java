class Solution {
    public boolean validPath(int n,
                             int[][] edges,
                             int source,
                             int destination) {

        List<List<Integer>> adjList = new ArrayList<>();

        for(int i = 0; i < n; i++) {
            adjList.add(new ArrayList<>());
        }

        for(int[] edge : edges) {

            int u = edge[0];
            int v = edge[1];

            adjList.get(u).add(v);
            adjList.get(v).add(u);
        }

        boolean[] visited = new boolean[n];

        return dfsHelper(
                adjList,
                visited,
                source,
                destination
        );
    }

    boolean dfsHelper(List<List<Integer>> adjList,
                      boolean[] visited,
                      int vertex,
                      int destination) {

        if(vertex == destination) {
            return true;
        }

        visited[vertex] = true;

        for(int neighbour : adjList.get(vertex)) {

            if(!visited[neighbour]) {

                if(dfsHelper(
                        adjList,
                        visited,
                        neighbour,
                        destination)) {

                    return true;
                }
            }
        }

        return false;
    }
}