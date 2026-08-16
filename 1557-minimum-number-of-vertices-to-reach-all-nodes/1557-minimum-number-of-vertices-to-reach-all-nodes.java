import java.util.*;

class Solution {
    public List<Integer> findSmallestSetOfVertices(int n, List<List<Integer>> edges) {
        boolean[] hasIncoming = new boolean[n];

        // Mark nodes that have an incoming edge
        for (List<Integer> edge : edges) {
            int to = edge.get(1);
            hasIncoming[to] = true;
        }

        // Nodes with no incoming edge are required starting vertices
        List<Integer> result = new ArrayList<>();

        for (int i = 0; i < n; i++) {
            if (!hasIncoming[i]) {
                result.add(i);
            }
        }

        return result;
    }
}