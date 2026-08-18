class Solution {
    public int[] scoreValidator(String[] events) {
        int n = events.length;
        int score = 0;
        int counter = 0;
        for(int i=0; i<n; i++) {
            if(counter == 10) break;
            String s = events[i];
            if(s.equals("0")) score += 0;
            if(s.equals("1")) score += 1;
            if(s.equals("2")) score += 2;
            if(s.equals("3")) score += 3;
            if(s.equals("4")) score += 4;
            if(s.equals("6")) score += 6;
            if(s.equals("W")) counter += 1;
            if(s.equals("WD")) score += 1;
            if(s.equals("NB")) score += 1;
        }
        return new int[]{score, counter};
    }
}