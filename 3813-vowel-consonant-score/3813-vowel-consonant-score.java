class Solution {
    public int vowelConsonantScore(String s) {
        int vowels = 0;
        int consonants = 0;

        for(int i =0; i <s.length();i++){
            char ch = s.charAt(i);

            if(Character.isLetter(ch)){
                if(ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'){
                    vowels++;
                }else{
                    consonants++;
                }
            }
        }
        if(consonants == 0){
            return 0;
        }
        return vowels/ consonants;
    }
}