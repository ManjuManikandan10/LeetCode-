class Solution {
    public int distributeCandies(int[] candyType) {
        int max = candyType.length / 2;
        HashSet<Integer> h = new HashSet<>();
        for(int x : candyType){
            h.add(x);
        }
        if(h.size()>=max){
            return max;
        }
    return h.size();
    }
}