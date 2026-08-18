class Solution {
    public int firstMatchingIndex(String s) {
        int end = s.length();
        for(int i=0;i<end;i++)
            if(s.charAt(i) == s.charAt(end-i-1)){
                return i;
            }
        return -1;
    }
}