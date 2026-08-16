import java.util.*;

class Solution {
    public int largestComponentSize(int[] nums) {
        int max = 100001;
        DSU dsu = new DSU(max);

        for (int num : nums) {
            // Factorization
            for (int i = 2; i * i <= num; i++) {
                if (num % i == 0) {
                    dsu.union(num, i);
                    dsu.union(num, num / i);
                }
            }
        }

        // Count components
        Map<Integer, Integer> count = new HashMap<>();
        int maxSize = 0;

        for (int num : nums) {
            int root = dsu.find(num);
            count.put(root, count.getOrDefault(root, 0) + 1);
            maxSize = Math.max(maxSize, count.get(root));
        }

        return maxSize;
    }
}

class DSU {
    int[] parent;

    public DSU(int size) {
        parent = new int[size];
        for (int i = 0; i < size; i++) {
            parent[i] = i;
        }
    }

    public int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]); // path compression
        }
        return parent[x];
    }

    public void union(int a, int b) {
        int pa = find(a);
        int pb = find(b);
        if (pa != pb) {
            parent[pa] = pb;
        }
    }
}