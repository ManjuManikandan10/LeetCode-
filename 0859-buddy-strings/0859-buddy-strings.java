class Solution {
    public boolean buddyStrings(String s, String goal) {
        if(s.length() != goal.length()){
            return false;
        }

        // for aa,aa or ab,ab
        if(s.equals(goal)){
            HashSet<Character> set = new HashSet<>();
            for(char c : s.toCharArray()){
                if(set.contains(c)){
                    return true;
                }
                else{
                    set.add(c);
                }
            }
            return false;
        }

        // actual code
        int pos1=-1;
        int pos2=-1;
        int mis = 0;
        for(int i =0;i<s.length();i++){
            if(s.charAt(i) != goal.charAt(i)){
                if(mis == 0 ){
                    pos1 = i;
                    mis++;
                }
                else if(mis == 1){
                    pos2 = i;
                    mis++;
                }
                else{
                    return false;
                }
            }
        }

        if(mis == 2){
            if(s.charAt(pos1) == goal.charAt(pos2) && s.charAt(pos2) == goal.charAt(pos1)){
                return true;
            }
            else{
                return false;
            }
        }

        return false;

    }
}