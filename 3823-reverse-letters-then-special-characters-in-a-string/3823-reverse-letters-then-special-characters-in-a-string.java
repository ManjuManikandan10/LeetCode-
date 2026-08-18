class Solution {
    public String reverseByType(String s) {
        boolean b[] = new boolean[s.length()]; 
        StringBuilder sb1 = new StringBuilder();
        StringBuilder sb2 = new StringBuilder();
        for(int i=0; i<s.length(); i++){
            if(s.charAt(i)>='a' && s.charAt(i)<='z'){
            sb1.append(s.charAt(i));
            b[i] = true;
            }else{
            sb2.append(s.charAt(i));
            b[i] = false;
            }
        }
        // sb1 = sb1.reverse();
        // sb2 = sb2.reverse();
        StringBuilder result = new StringBuilder();
        int l1 = sb1.length()-1;
        int l2 = sb2.length()-1;
        for(int i=0; i<b.length; i++){
            if(b[i]==true){
            result.append(sb1.charAt(l1));
            l1--;
            }else{
            result.append(sb2.charAt(l2));
            l2--;
            }
        }
        //System.out.println(result);
        return result.toString();
    }
}