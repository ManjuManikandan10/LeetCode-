class Solution {
    public int residuePrefixes(String s) {
        int count = 0;
        HashSet<Character> set = new HashSet<>();
        for(int i = 0; i < s.length(); i++){
            set.add(s.charAt(i));
            if(set.size() == (i + 1) % 3){
                count++;
            }
        }
        return count;
    }
}