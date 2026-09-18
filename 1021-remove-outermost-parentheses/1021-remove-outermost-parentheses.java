class Solution {
    public String removeOuterParentheses(String s) {
        int sp = -1;
        StringBuilder sb = new StringBuilder(s.length());
        for(char c : s.toCharArray()) {
            if(c == '(') {
                sp++;
                if(sp >= 1) sb.append(c);
                continue;
            }
            sp--;
            if(sp > -1) sb.append(c);
        }
        return sb.toString();
    }
}