class Solution {
    public String reversePrefix(String s, int k) {
        StringBuilder r = new StringBuilder(s.substring(0, k));
        r.reverse();
        r.append(s.substring(k, s.length()));
        return r.toString();
    }
}