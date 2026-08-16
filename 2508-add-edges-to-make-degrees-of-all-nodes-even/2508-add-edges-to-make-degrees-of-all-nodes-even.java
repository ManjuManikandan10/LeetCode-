class Solution {
    public boolean isPossible(int m, List<List<Integer>> Edges) {
        Map<Integer, Set<Integer>> map = new HashMap<>(m);
        for (int i = 1; i <= m; i++) map.put(i,new HashSet<>());

        for (List<Integer> items : Edges) {
            map.get(items.get(0)).add(items.get(1));
            map.get(items.get(1)).add(items.get(0));
        }

        List<Integer> odd = new ArrayList<>(4);
        for (int val : map.keySet()) if ((map.get(val).size() & 1) == 1) odd.add(val);

        if (odd.isEmpty()) return true;
        if (odd.size() == 2) {
            int a = odd.get(0); int b = odd.get(1);
            if (!map.get(a).contains(b)) return true;
            else {
                for (int key : map.keySet()) {
                    if ((map.get(key).size() & 1) == 0 && !map.get(key).contains(a) && !map.get(key).contains(b))
                        return true;
                }
            }
        }
        if (odd.size() == 4) {
            int a = odd.get(0); int b = odd.get(1);
            int c = odd.get(2); int d = odd.get(3);

            if (!map.get(a).contains(b) && !map.get(c).contains(d)) return true;
            if (!map.get(c).contains(a) && !map.get(b).contains(d)) return true;
            return !map.get(a).contains(d) && !map.get(c).contains(b);
        }

        return false;
    }
}