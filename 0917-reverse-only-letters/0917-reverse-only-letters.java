class Solution {
    public String reverseOnlyLetters(String s) {
        int i=0;
        int j=s.length()-1;
        char[] c=s.toCharArray();
        while(i<j){
            if(!Character.isLetter(c[i])){
                i++;
                continue;
            }
            if(!Character.isLetter(c[j])){
                j--;
                continue;
            }
            char t=c[i];
            c[i]=c[j];
            c[j]=t;
            i++;
            j--;
        }
        return new String(c);
    }
}