class Solution {
    public int firstUniqueEven(int[] nums) {
        ArrayList<Integer> arr = new ArrayList<>();
        Set<Integer> set = new HashSet<>();

        for (int i : nums) {
            if (i % 2 == 1)
                continue;

            if (set.contains(i))
                continue;

            else if (arr.contains(i)) {
                set.add(i);
                arr.remove((Integer) i);
            } else
                arr.add(i);
        }

        if (arr.isEmpty())
            return -1;

        return arr.get(0);
    }
}